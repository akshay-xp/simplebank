package grpc_api

import (
	"fmt"

	db "github.com/akshay-xp/simplebank/db/sqlc"
	"github.com/akshay-xp/simplebank/pb"
	"github.com/akshay-xp/simplebank/token"
	"github.com/akshay-xp/simplebank/util"
)

// Server serves gRPC requests for out banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer create a new gRPC server.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}

	return server, nil
}
