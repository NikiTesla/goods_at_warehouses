package jsonrpc

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/NikiTesla/goods_at_warehouses"
	"github.com/NikiTesla/goods_at_warehouses/pkg/database"
)

// Warehouses type is struct to interact with warehouses, has environment with bd as a field
type Warehouses struct {
	db database.DataBase
}

// Create gets list of Warehouses and ask database to create them
// puts in reply list of successfully created
func (wH *Warehouses) Create(args []goods_at_warehouses.Warehouse, reply *[]goods_at_warehouses.Warehouse) error {
	created := make([]goods_at_warehouses.Warehouse, 0, len(args))
	for _, warehouse := range args {
		if err := wH.db.CreateWarehouse(warehouse); err != nil {
			log.WithError(err).Error("cannot create warehouse")
			continue
		}
		created = append(created, warehouse)
	}
	*reply = created
	return nil
}

// GetAmount map with good_code and warehouse_id fields, puts in reply amount of available goods at the warehouse
func (wH *Warehouses) GetAmount(args map[string]int, reply *int) error {
	goodCode, ok := args["goodCode"]
	if !ok {
		return fmt.Errorf("request is incorrect, good_code is not presented")
	}
	warehouseID, ok := args["warehouseID"]
	if !ok {
		return fmt.Errorf("request is incorrect, warehouse_id is not presented")
	}
	amount, err := wH.db.GetAmount(goodCode, warehouseID)
	if err == nil {
		*reply = amount
	}
	return err
}
