package database

import (
	"fmt"
	"log"

	lamodatest "github.com/NikiTesla/lamoda_test"
)

func (db *PostgresDB) CreateGood(good lamodatest.Good) error {
	fmt.Printf("Creating good %v", good)
	query := "INSERT INTO goods(name, code, size, amount) VALUES ($1, $2, $3, $4)"

	var exists bool
	row := db.DB.QueryRow("SELECT EXISTS(SELECT id FROM goods WHERE code = $1)",
		good.Code)
	row.Scan(&exists)

	if exists {
		return fmt.Errorf("good already exists")
	}

	_, err := db.DB.Exec(query, good.Name, good.Code, good.Size, good.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (db *PostgresDB) AddGood(good_code, warehouse_id, amount int) error {
	return db.doIfAvailable(warehouse_id, func() error {
		var available_amount int
		row := db.DB.QueryRow("SELECT amount FROM goods WHERE code = $1", good_code)
		row.Scan(&available_amount)
		if available_amount < amount {
			return fmt.Errorf("there is not enough %d good in %d warehouse. Available only %d",
				good_code, warehouse_id, available_amount)
		}

		_, err := db.DB.Exec("UPDATE goods SET amount = amount - $1 WHERE code = $2", amount, good_code)
		if err != nil {
			return err
		}

		var exists bool
		row = db.DB.QueryRow("SELECT EXISTS(SELECT id FROM warehouse_goods WHERE good_code = $1 AND warehouse_id = $2)",
			good_code, warehouse_id)
		row.Scan(&exists)

		var query string
		if !exists {
			query = "INSERT INTO warehouse_goods(good_code, warehouse_id, available_amount) VALUES ($1, $2, $3)"
		} else {
			query = "UPDATE warehouse_goods SET available_amount = available_amount + $3 WHERE good_code = $1 AND warehouse_id = $2"
		}

		if _, err := db.DB.Exec(query, good_code, warehouse_id, amount); err != nil {
			return err
		}

		return nil
	})
}

func (db *PostgresDB) ReserveGood(good_code, warehouse_id, amount int) error {
	return db.doIfAvailable(warehouse_id, func() error {
		log.Printf("Reserving %d at the warehouse %d in amount of %d", good_code, warehouse_id, amount)
		var available_amount int
		row := db.DB.QueryRow("SELECT available_amount FROM warehouse_goods WHERE good_code = $1 AND warehouse_id = $2",
			good_code, warehouse_id)
		row.Scan(&available_amount)

		if available_amount < amount {
			return fmt.Errorf("there is not enough %d good in warehouse %d", good_code, warehouse_id)
		}

		query := "UPDATE warehouse_goods SET " +
			"available_amount = available_amount - $2, reserved_amount = reserved_amount + $2 " +
			"WHERE good_code = $1 AND warehouse_id = $3"

		if _, err := db.DB.Exec(query, good_code, amount, warehouse_id); err != nil {
			return err
		}
		return nil
	})
}

func (db *PostgresDB) CancelGoodReservation(good_code, warehouse_id, amount int) error {
	return db.doIfAvailable(warehouse_id, func() error {
		log.Printf("Canceling reservation %d at the warehouse %d in amount of %d", good_code, warehouse_id, amount)
		var reserved_amount int
		row := db.DB.QueryRow("SELECT reserved_amount FROM warehouse_goods WHERE good_code = $1 AND warehouse_id = $2",
			good_code, warehouse_id)
		row.Scan(&reserved_amount)

		if reserved_amount < amount {
			return fmt.Errorf("there is not enough reserved goods %d in warehouse %d", good_code, warehouse_id)
		}

		query := "UPDATE warehouse_goods SET " +
			"available_amount = available_amount + $2, reserved_amount = reserved_amount - $2 " +
			"WHERE good_code = $1 AND warehouse_id = $3"

		_, err := db.DB.Exec(query, good_code, amount, warehouse_id)
		if err != nil {
			return err
		}

		return nil
	})
}

func (db *PostgresDB) doIfAvailable(warehouse_id int, f func() error) error {
	var err error
	if _, err = db.DB.Exec("BEGIN;"); err != nil {
		return err
	}

	var availability bool
	row := db.DB.QueryRow("SELECT availability FROM warehouses WHERE id = $1", warehouse_id)
	row.Scan(&availability)
	if !availability {
		return fmt.Errorf("warehouse not available at the moment")
	}

	if err = f(); err != nil {
		return err
	}

	row = db.DB.QueryRow("SELECT availability FROM warehouses WHERE id = $1", warehouse_id)
	row.Scan(&availability)
	if !availability {
		db.DB.Exec("ROLLBACK;")
		return fmt.Errorf("warehouse became unavailable while adding goods")
	}

	if _, err := db.DB.Exec("COMMIT;"); err != nil {
		return err
	}

	return nil
}
