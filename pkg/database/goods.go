package database

import (
	"fmt"
	"log"

	lamodatest "github.com/NikiTesla/lamoda_test"
	"github.com/NikiTesla/lamoda_test/pkg/environment"
)

func CreateGood(good *lamodatest.Good, env *environment.Environment) error {
	fmt.Printf("Creating good %v", good)
	query := "INSERT INTO goods(name, code, size, amount) VALUES ($1, $2, $3, $4)"

	_, err := env.DB.Exec(query, good.Name, good.Code, good.Size, good.Amount)
	if err != nil {
		return err
	}
	return nil
}

func AddGood(good_code, warehouse_id, amount int, env *environment.Environment) error {
	return doIfAvailable(warehouse_id, env, func() error {
		var available_amount int
		row := env.DB.QueryRow("SELECT amount FROM goods WHERE code = $1", good_code)
		row.Scan(&available_amount)
		if available_amount < amount {
			return fmt.Errorf("there is not enough %d good in %d warehouse. Available only %d",
				good_code, warehouse_id, available_amount)
		}

		_, err := env.DB.Exec("UPDATE goods SET amount = amount - $1 WHERE code = $2", amount, good_code)
		if err != nil {
			return err
		}

		var exists bool
		row = env.DB.QueryRow("SELECT EXISTS(SELECT id FROM warehouse_goods WHERE good_code = $1 AND warehouse_id = $2)",
			good_code, warehouse_id)
		row.Scan(&exists)

		var query string
		if !exists {
			query = "INSERT INTO warehouse_goods(good_code, warehouse_id, available_amount) VALUES ($1, $2, $3)"
		} else {
			query = "UPDATE warehouse_goods SET available_amount = available_amount + $3 WHERE good_code = $1 AND warehouse_id = $2"
		}

		if _, err := env.DB.Exec(query, good_code, warehouse_id, amount); err != nil {
			return err
		}

		return nil
	})
}

func ReserveGood(good_code, warehouse_id, amount int, env *environment.Environment) error {
	return doIfAvailable(warehouse_id, env, func() error {
		log.Printf("Reserving %d on warehouse %d in amount of %d", good_code, warehouse_id, amount)
		var available_amount int
		row := env.DB.QueryRow("SELECT available_amount FROM warehouse_goods WHERE good_code = $1 AND warehouse_id = $2)",
			good_code, warehouse_id)
		row.Scan(&available_amount)

		if available_amount < amount {
			return fmt.Errorf("there is not enough %d good in warehouse %d", good_code, warehouse_id)
		}

		query := "UPDATE warehouse_goods SET " +
			"available_amount = available_amount - $2, reserved_amount = reserved_amount + $2 " +
			"WHERE good_code = $1 AND warehouse_id = $3"

		if _, err := env.DB.Exec(query, good_code, amount, warehouse_id); err != nil {
			return err
		}
		return nil
	})
}

func CancelGoodReservation(good_code, warehouse_id, amount int, env *environment.Environment) error {
	return doIfAvailable(warehouse_id, env, func() error {
		log.Printf("Canceling reservation %d on warehouse %d in amount of %d", good_code, warehouse_id, amount)
		var available_amount int
		row := env.DB.QueryRow("SELECT available_amount FROM warehouse_goods WHERE good_code = $1 AND warehouse_id = $2)",
			good_code, warehouse_id)
		row.Scan(&available_amount)

		if available_amount < amount {
			return fmt.Errorf("there is not enough %d good in warehouse %d", good_code, warehouse_id)
		}

		query := "UPDATE warehouse_goods SET " +
			"available_amount = available_amount + $2, reserved_amount = reserved_amount - $2 " +
			"WHERE good_code = $1 AND warehouse_id = $3"

		_, err := env.DB.Exec(query, good_code, amount, warehouse_id)
		if err != nil {
			return err
		}

		return nil
	})
}

func doIfAvailable(warehouse_id int, env *environment.Environment, f func() error) error {
	var err error
	if _, err = env.DB.Exec("START TRANSACTION"); err != nil {
		return err
	}

	var availability bool
	row := env.DB.QueryRow("SELECT availability FROM warehouses WHERE id = $1", warehouse_id)
	row.Scan(&availability)
	if !availability {
		return fmt.Errorf("warehouse not available at the moment")
	}

	if err = f(); err != nil {
		return err
	}

	row = env.DB.QueryRow("SELECT availability FROM warehouses WHERE id = $1", warehouse_id)
	row.Scan(&availability)
	if !availability {
		env.DB.Exec("ROLLBACK")
		return fmt.Errorf("warehouse became unavailable while adding goods")
	}

	if _, err := env.DB.Exec("COMMIT"); err != nil {
		return err
	}

	return nil
}
