package database

import (
	"fmt"

	lamodatest "github.com/NikiTesla/lamoda_test"
)

// CreateGood gets Good, check if it exists.
// If does not exist - insert it into goods
func (db *PostgresDB) CreateGood(good lamodatest.Good) error {
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

// AddGood gets goodCode, warehouseID, amount of goods to be added at the warehouse.
// Checks if there are enough goods at the sort center (joint storage)
// If there are such goods at the warehouse, sum their amounts. If not - inserts new row in the table
func (db *PostgresDB) AddGood(goodCode, warehouseID, amount int) error {
	return db.doIfAvailable(warehouseID, func() error {
		var availableAmount int
		row := db.DB.QueryRow("SELECT amount FROM goods WHERE code = $1", goodCode)
		if err := row.Scan(&availableAmount); err != nil {
			return fmt.Errorf("error occured while scanning row, error: %s", err)
		}

		if availableAmount < amount {
			return fmt.Errorf("there is not enough %d good in %d warehouse. Available only %d",
				goodCode, warehouseID, availableAmount)
		}

		_, err := db.DB.Exec("UPDATE goods SET amount = amount - $1 WHERE code = $2", amount, goodCode)
		if err != nil {
			return err
		}

		var exists bool
		row = db.DB.QueryRow("SELECT EXISTS(SELECT id FROM warehouse_goods WHERE good_code = $1 AND warehouse_id = $2)",
			goodCode, warehouseID)
		row.Scan(&exists)

		var query string
		if !exists {
			query = "INSERT INTO warehouse_goods(good_code, warehouse_id, available_amount) VALUES ($1, $2, $3)"
		} else {
			query = "UPDATE warehouse_goods SET available_amount = available_amount + $3 WHERE good_code = $1 AND warehouse_id = $2"
		}

		if _, err := db.DB.Exec(query, goodCode, warehouseID, amount); err != nil {
			return err
		}

		return nil
	})
}

// ReserveGood uses doIfAvailable to reserve goods safely
// Gets good code and warehouse id, check if there are enough available goods at the warehouse
func (db *PostgresDB) ReserveGood(goodCode, warehouseID, amount int) error {
	return db.doIfAvailable(warehouseID, func() error {
		var availableAmount int
		row := db.DB.QueryRow("SELECT available_amount FROM warehouse_goods WHERE good_code = $1 AND warehouse_id = $2",
			goodCode, warehouseID)
		row.Scan(&availableAmount)

		if availableAmount < amount {
			return fmt.Errorf("there is not enough %d good in warehouse %d. Available only %d",
				goodCode, warehouseID, availableAmount)
		}

		query := "UPDATE warehouse_goods SET " +
			"available_amount = available_amount - $2, reserved_amount = reserved_amount + $2 " +
			"WHERE good_code = $1 AND warehouse_id = $3"

		if _, err := db.DB.Exec(query, goodCode, amount, warehouseID); err != nil {
			return err
		}
		return nil
	})
}

// CancelGoodReservation uses doIfAvailable to cancel reservation of good safely
// Gets good code and warehouse id, check if there are enough reserved goods at the warehouse
func (db *PostgresDB) CancelGoodReservation(goodCode, warehouseID, amount int) error {
	return db.doIfAvailable(warehouseID, func() error {
		var reservedAmount int
		row := db.DB.QueryRow("SELECT reserved_amount FROM warehouse_goods WHERE good_code = $1 AND warehouse_id = $2",
			goodCode, warehouseID)
		if err := row.Scan(&reservedAmount); err != nil {
			return fmt.Errorf("error occured while scanning row, error: %s", err)
		}

		if reservedAmount < amount {
			return fmt.Errorf("there is not enough reserved goods %d at the warehouse %d. Reserved only %d",
				goodCode, warehouseID, reservedAmount)
		}

		query := "UPDATE warehouse_goods SET " +
			"available_amount = available_amount + $2, reserved_amount = reserved_amount - $2 " +
			"WHERE good_code = $1 AND warehouse_id = $3"

		_, err := db.DB.Exec(query, goodCode, amount, warehouseID)
		if err != nil {
			return err
		}

		return nil
	})
}

// doIfAvailable is method to wrap functional in transaction shell
func (db *PostgresDB) doIfAvailable(warehouseID int, f func() error) error {
	var err error
	if _, err = db.DB.Exec("BEGIN;"); err != nil {
		return err
	}

	var availability bool
	row := db.DB.QueryRow("SELECT availability FROM warehouses WHERE id = $1", warehouseID)
	if err := row.Scan(&availability); err != nil {
		return fmt.Errorf("error occured while scanning row, error: %s", err)
	}

	if !availability {
		return fmt.Errorf("warehouse not available at the moment")
	}

	if err = f(); err != nil {
		return err
	}

	row = db.DB.QueryRow("SELECT availability FROM warehouses WHERE id = $1", warehouseID)
	if err := row.Scan(&availability); err != nil {
		return fmt.Errorf("error occured while scanning row, error: %s", err)
	}

	if !availability {
		db.DB.Exec("ROLLBACK;")
		return fmt.Errorf("warehouse became unavailable while adding goods")
	}

	if _, err := db.DB.Exec("COMMIT;"); err != nil {
		return err
	}

	return nil
}
