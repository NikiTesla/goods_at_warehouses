package database

import (
	"fmt"
	"log"

	"github.com/NikiTesla/goods_at_warehouses/pkg/core"
)

// CreateWarehouse check if warehouse exist
// If not - inserts into warehouses table Warehouse.Name and Warehouse.Availability
func (db *PostgresDB) CreateWarehouse(warehouse core.Warehouse) error {
	var exists bool
	err := db.DB.QueryRow("SELECT EXISTS(SELECT id FROM warehouses WHERE name = $1)",
		warehouse.Name).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("warehouse already exists")
	}

	_, err = db.DB.Exec("INSERT INTO warehouses(name, availability) VALUES ($1, $2)",
		warehouse.Name, warehouse.Availability)

	return err
}

// GetAmount gets goodCode and warehouseID, returns available amount of goods at the warehouse
func (db *PostgresDB) GetAmount(goodCode, warehouseID int) (int, error) {
	var amount int

	query := "SELECT available_amount FROM warehouse_goods WHERE good_code = $1 AND warehouse_id = $2"
	err := db.DB.QueryRow(query, goodCode, warehouseID).Scan(&amount)
	if err != nil {
		log.Println(err)
		return 0, fmt.Errorf("there is no %d goods at the %d warehouse", goodCode, warehouseID)
	}
	return amount, nil
}
