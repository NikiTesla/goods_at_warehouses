package jsonrpc

import (
	"fmt"
	"log"

	lamodatest "github.com/NikiTesla/lamoda_test"
	"github.com/NikiTesla/lamoda_test/pkg/database"
	"github.com/NikiTesla/lamoda_test/pkg/environment"
)

type Goods struct {
	env *environment.Environment
}

func (g *Goods) Create(args []lamodatest.Good, reply *Reply) error {
	for _, good := range args {
		log.Printf("Creating good: %v", good)
		if err := database.CreateGood(&good, g.env); err != nil {
			log.Print("Error while creating goods ", err.Error())
			return err
		}
	}
	*reply = Reply{fmt.Sprintf("Created goods %v: ", args)}

	return nil
}

func (g *Goods) Add(args []map[string]int, reply *Reply) error {
	log.Printf("Adding goods: %v", args)
	for _, arg := range args {
		good_code, warehouse_id, amount, err := parseGoods(arg)
		if err != nil {
			log.Print("Error occured while parsing args, error: ", err.Error())
			return err
		}
		if err := database.AddGood(good_code, warehouse_id, amount, g.env); err != nil {
			log.Print("Cannot add good, error: ", err.Error())
			return err
		}
	}
	*reply = Reply{fmt.Sprintf("Added goods %v", args)}

	return nil
}

func (g *Goods) Reserve(args []map[string]int, reply *Reply) error {
	log.Print("Reserving: ", args)
	for _, arg := range args {
		good_code, warehouse_id, amount, err := parseGoods(arg)
		if err != nil {
			log.Print("Error while parsing args ", err.Error())
			return err
		}

		if err = database.ReserveGood(good_code, warehouse_id, amount, g.env); err != nil {
			log.Printf("error occured while reserving good with code %d, error: %s", good_code, err.Error())
			return err
		}
	}
	*reply = Reply{fmt.Sprintf("Reserved goods: %v", args)}

	log.Print("reserved")

	return nil
}

func (g *Goods) CancelReservation(args []map[string]int, reply *Reply) error {
	log.Print("Canceling reservations: ", args)
	for _, arg := range args {
		good_code, warehouse_id, amount, err := parseGoods(arg)
		if err != nil {
			log.Print("Error while parsing args ", err.Error())
			return err
		}

		if err := database.CancelGoodReservation(good_code, warehouse_id, amount, g.env); err != nil {
			log.Printf("error occured while cancelling reservation of good with code %d, error: %s", good_code, err.Error())
			return err
		}
	}

	*reply = Reply{fmt.Sprintf("Cancel reservations of goods: %v", args)}

	log.Print("cancelled")

	return nil
}

func parseGoods(arg map[string]int) (int, int, int, error) {
	good_code, ok := arg["good_code"]
	if !ok {
		return 0, 0, 0, fmt.Errorf("incorrect request, good_code not presented")
	}
	warehouse_id, ok := arg["warehouse_id"]
	if !ok {
		return 0, 0, 0, fmt.Errorf("incorrect request, warehouse_id not presented")
	}
	amount, ok := arg["amount"]
	if !ok {
		return 0, 0, 0, fmt.Errorf("incorrect request, amount not presented")
	}
	return good_code, warehouse_id, amount, nil
}
