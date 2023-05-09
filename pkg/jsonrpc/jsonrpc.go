package jsonrpc

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/NikiTesla/lamoda_test/pkg/environment"
)

type Server struct {
	rpcServer *rpc.Server
	env       *environment.Environment
}

type Reply struct {
	Data string
}

func NewServer(env *environment.Environment) *Server {
	rpcServer := rpc.NewServer()

	if err := rpc.RegisterName("Goods", &Goods{env}); err != nil {
		log.Println("cannot register goods: error", err.Error())
	}
	if err := rpc.RegisterName("Warehouses", &Warehouses{env}); err != nil {
		log.Println("cannot register goods: error", err.Error())
	}

	return &Server{
		rpcServer: rpcServer,
		env:       env,
	}
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
