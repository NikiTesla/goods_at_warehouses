package jsonrpc

import (
	"fmt"
	"log"

	lamodatest "github.com/NikiTesla/lamoda_test"
	"github.com/NikiTesla/lamoda_test/pkg/database"
)

// Warehouses type is struct to interact with warehouses, has environment with bd as a field
type Warehouses struct {
	db database.DataBase
}

// Create gets list of Warehouses and ask database to create them
// puts in reply list of successfully created
func (wH *Warehouses) Create(args []lamodatest.Warehouse, reply *Reply) error {
	created := make([]lamodatest.Warehouse, 0, len(args))
	for _, warehouse := range args {
		log.Printf("Creating warehouse %v", warehouse)

		if err := wH.db.CreateWarehouse(warehouse); err != nil {
			log.Print("Error occured while creating warehouse ", err)
			continue
		}
		created = append(created, warehouse)
	}

	*reply = Reply{fmt.Sprintf("Created warehouses: %v", created)}

	return nil
}

// GetAmount map with good_code and warehouse_id fields, puts in reply amount of available goods at the warehouse
func (wH *Warehouses) GetAmount(args map[string]int, reply *Reply) error {
	goodCode, ok := args["goodCode"]
	if !ok {
		return fmt.Errorf("request is incorrect, good_code is not presented")
	}
	warehouseID, ok := args["warehouseID"]
	if !ok {
		return fmt.Errorf("request is incorrect, warehouse_id is not presented")
	}

	log.Printf("Getting amount of %d in warehouse %d", goodCode, warehouseID)
	amount, err := wH.db.GetAmount(goodCode, warehouseID)
	if err == nil {
		*reply = Reply{fmt.Sprint("amount: ", amount)}
	}

	return err
}
