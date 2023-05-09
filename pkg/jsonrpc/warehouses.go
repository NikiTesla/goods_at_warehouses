package jsonrpc

import (
	"fmt"
	"log"

	lamodatest "github.com/NikiTesla/lamoda_test"
	"github.com/NikiTesla/lamoda_test/pkg/database"
	"github.com/NikiTesla/lamoda_test/pkg/environment"
)

type Warehouses struct {
	env *environment.Environment
}

func (wH *Warehouses) Create(args []lamodatest.Warehouse, reply *Reply) error {
	for _, warehouse := range args {
		log.Printf("Creating warehouse %v", warehouse)
		if err := database.CreateWarehouse(warehouse, wH.env); err != nil {
			log.Print("Error occured while creating warehouse ", err)
			return err
		}
	}

	*reply = Reply{fmt.Sprintf("Created warehouses: %v", args)}

	return nil
}

func (wH *Warehouses) GetAmount(args map[string]int, reply *Reply) error {
	good_code, ok := args["good_code"]
	if !ok {
		return fmt.Errorf("request is incorrect, good_code is not presented")
	}
	warehouse_id, ok := args["warehouse_id"]
	if !ok {
		return fmt.Errorf("request is incorrect, warehouse_id is not presented")
	}

	log.Printf("Getting amount of %d in warehouse %d", good_code, warehouse_id)
	amount, err := database.GetAmount(good_code, warehouse_id, wH.env)
	if err == nil {
		*reply = Reply{fmt.Sprint("amount: ", amount)}
	}

	return err
}
