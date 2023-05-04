package jsonrpc

import "fmt"

type Goods int

func (g *Goods) Add(args []string, reply *Reply) error {
	for _, arg := range args {
		fmt.Println("added good", arg)
	}
	*reply = Reply{"Nice try"}

	return nil
}
