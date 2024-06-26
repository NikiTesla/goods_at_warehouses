package jsonrpc

import (
	log "github.com/sirupsen/logrus"

	"github.com/NikiTesla/goods_at_warehouses/pkg/core"
	"github.com/NikiTesla/goods_at_warehouses/pkg/database"
)

const (
	statusCreated              = "created"
	statusAdded                = "added"
	statusReserved             = "reserved"
	statusReservationCancelled = "reservation cancelled"
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
func (g *Goods) Create(args []core.Good, reply *[]core.Good) error {
	created := make([]core.Good, 0, len(args))
	for _, good := range args {
		if err := g.db.CreateGood(good); err != nil {
			log.WithError(err).Errorf("cannot create good with code %d", good.Code)
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
			log.WithError(err).Errorf("cannot add good with code %d", arg.GoodCode)
			continue
		}
		arg.Status = statusAdded
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
			log.WithError(err).Errorf("cannot reserve good with code %d", arg.GoodCode)
			continue
		}
		arg.Status = statusReserved
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
			log.WithError(err).Printf("error occured while cancelling reservation of good with code %d, error: %s\n", arg.GoodCode, err.Error())
			continue
		}
		arg.Status = statusReservationCancelled
		cancelled = append(cancelled, arg)
	}
	*reply = cancelled
	return nil
}
