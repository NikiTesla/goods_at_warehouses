package jsonrpc

import (
	"log"

	lamodatest "github.com/NikiTesla/lamoda_test"
	"github.com/NikiTesla/lamoda_test/pkg/database"
)

// Goods is struct to interact with goods at different warehouses, has environment with bd as a field
type Goods struct {
	db database.DataBase
}

// WarehouseGoods is a structure to store information about the actions with goods at warehouse
type WarehouseGoodAction struct {
	GoodCode    int    `json:"goodCode"`
	WarehouseID int    `json:"warehouseID"`
	Amount      int    `json:"amount"`
	Status      string `json:"status"` // added, reserved or reservation cancelled
}

// Create gets list of goods as []Good, ask database to create them
// put in reply successfully created goods
func (g *Goods) Create(args []lamodatest.Good, reply *[]lamodatest.Good) error {
	created := make([]lamodatest.Good, 0, len(args))
	for _, good := range args {
		if err := g.db.CreateGood(good); err != nil {
			log.Printf("Error while creating good %v, error: %s\n", good, err.Error())
			continue
		}
		created = append(created, good)
	}
	*reply = created

	return nil
}

// Add gets list of maps with goodCode, warehouseID and amount of goods to be added
// put in reply successfully added goods
func (g *Goods) Add(args []WarehouseGoodAction, reply *[]WarehouseGoodAction) error {
	added := make([]WarehouseGoodAction, 0, len(args))
	for _, arg := range args {
		if err := g.db.AddGood(arg.GoodCode, arg.WarehouseID, arg.Amount); err != nil {
			log.Printf("Cannot add good with code %d, error: %s\n", arg.GoodCode, err.Error())
			continue
		}
		arg.Status = "added"
		added = append(added, arg)
	}
	*reply = added

	return nil
}

// Reserve gets list of maps with goodCode, warehouseID and amount of goods to be reserved
// put in reply successfully reserved goods
func (g *Goods) Reserve(args []WarehouseGoodAction, reply *[]WarehouseGoodAction) error {
	reserved := make([]WarehouseGoodAction, 0, len(args))
	for _, arg := range args {
		if err := g.db.ReserveGood(arg.GoodCode, arg.WarehouseID, arg.Amount); err != nil {
			log.Printf("error occured while reserving good with code %d, error: %s\n", arg.GoodCode, err.Error())
			continue
		}
		arg.Status = "reserved"
		reserved = append(reserved, arg)
	}

	*reply = reserved

	return nil
}

// CancelReservation gets lsit of maps with goodCode, warehouseID and amount of goods to cancel reservation
// put in reply successfully cancelled reservations of goods
func (g *Goods) CancelReservation(args []WarehouseGoodAction, reply *[]WarehouseGoodAction) error {
	cancelled := make([]WarehouseGoodAction, 0, len(args))
	for _, arg := range args {
		if err := g.db.CancelGoodReservation(arg.GoodCode, arg.WarehouseID, arg.Amount); err != nil {
			log.Printf("error occured while cancelling reservation of good with code %d, error: %s\n", arg.GoodCode, err.Error())
			continue
		}
		arg.Status = "reservation cancelled"
		cancelled = append(cancelled, arg)
	}

	*reply = cancelled

	return nil
}
