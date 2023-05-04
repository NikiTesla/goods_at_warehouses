package jsonrpc

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Server struct {
	rpcServer *rpc.Server
}

type Reply struct {
	Data string
}

func NewServer() *Server {
	rpcServer := rpc.NewServer()

	if err := rpc.RegisterName("Goods", new(Goods)); err != nil {
		log.Println("cannot register goods: error", err.Error())
	}
	if err := rpc.RegisterName("Warehouses", new(Warehouses)); err != nil {
		log.Println("cannot register goods: error", err.Error())
	}

	return &Server{rpcServer: rpcServer}
}

func (s *Server) Run(port int) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("listen error: %s", err.Error())
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		jsonrpc.ServeConn(conn)
	}
}
