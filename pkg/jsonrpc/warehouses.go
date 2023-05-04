package jsonrpc

import (
	"fmt"
)

type Warehouses int

func (wH *Warehouses) Add(args *[]string, reply *Reply) error {
	for _, arg := range *args {
		fmt.Println("added warehouse", arg)
	}
	*reply = Reply{"Nice try"}

	return nil
}
