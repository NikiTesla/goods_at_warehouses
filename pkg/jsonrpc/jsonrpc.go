package jsonrpc

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	log "github.com/sirupsen/logrus"

	"github.com/NikiTesla/goods_at_warehouses/pkg/database"
	"github.com/NikiTesla/goods_at_warehouses/pkg/environment"
)

type Server struct {
	rpcServer *rpc.Server
	env       *environment.Environment
}

type Reply struct {
	Data string
}

// NewServer creates Server with rpc Server and environment as fields
// also registrates named methods
func NewServer(env *environment.Environment) *Server {
	rpcServer := rpc.NewServer()

	if err := rpc.RegisterName("Goods", &Goods{db: &database.PostgresDB{DB: env.DB}}); err != nil {
		log.WithError(err).Error("cannot register goods")
	}
	if err := rpc.RegisterName("Warehouses", &Warehouses{db: &database.PostgresDB{DB: env.DB}}); err != nil {
		log.WithError(err).Error("cannot register warehouses")
	}

	return &Server{
		rpcServer: rpcServer,
		env:       env,
	}
}

// Run runs server on port that is waiting for connections
// and run them in differnet goroutines
func (s *Server) Run(port int) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen port %d: %w", port, err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.WithError(err).Error("failed to accept connection")
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
