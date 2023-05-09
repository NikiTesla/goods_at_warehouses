package database

import (
	"fmt"

	lamodatest "github.com/NikiTesla/lamoda_test"
)

func (db *PostgresDB) CreateWarehouse(warehouse lamodatest.Warehouse) error {
	fmt.Printf("Creating warehouse %v", warehouse)

	var exists bool
	row := db.DB.QueryRow("SELECT EXISTS(SELECT id FROM warehouses WHERE name = $1)",
		warehouse.Name)
	row.Scan(&exists)

	if exists {
		return fmt.Errorf("warehouse already exists")
	}

	_, err := db.DB.Exec("INSERT INTO warehouses(name, availability) VALUES ($1, $2)",
		warehouse.Name, warehouse.Availability)

	return err
}

func (db *PostgresDB) GetAmount(good_code, warehouse_id int) (int, error) {
	fmt.Printf("Getting amount of %d", good_code)

	var amount int
	row := db.DB.QueryRow("SELECT available_amount FROM warehouse_goods WHERE good_code = $1 AND warehouse_id = $2",
		good_code, warehouse_id)
	if err := row.Scan(&amount); err != nil {
		return 0, fmt.Errorf("there is no %d good in %d warehouse", good_code, warehouse_id)
	}

	return amount, nil
}
