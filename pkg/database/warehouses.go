package database

import (
	"fmt"

	lamodatest "github.com/NikiTesla/lamoda_test"
	"github.com/NikiTesla/lamoda_test/pkg/environment"
)

func CreateWarehouse(warehouse lamodatest.Warehouse, env *environment.Environment) error {
	fmt.Printf("Creating good %v", warehouse)

	_, err := env.DB.Exec("INSERT INTO warehouses(name, availability) VALUES ($1, $2)",
		warehouse.Name, warehouse.Availability)

	return err
}

func GetAmount(good_code, warehouse_id int, env *environment.Environment) (int, error) {
	fmt.Printf("Getting amount of %d", good_code)

	var amount int
	row := env.DB.QueryRow("SELECT available_amount FROM warehouse_goods WHERE good_code = $1 AND warehouse_id = $2",
		good_code, warehouse_id)
	if err := row.Scan(&amount); err != nil {
		return 0, fmt.Errorf("there is no %d good in %d warehouse", good_code, warehouse_id)
	}

	return amount, nil
}
