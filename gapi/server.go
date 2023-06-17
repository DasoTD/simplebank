package gapi

import (
	"fmt"

	db "github.com/dasotd/simplebank/db/sqlc"
	"github.com/dasotd/simplebank/pb"
	"github.com/dasotd/simplebank/token"
	"github.com/dasotd/simplebank/util"
	// "github.com/dasotd/simplebank/worker"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	// taskDistributor worker.TaskDistributor
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		// taskDistributor: taskDistributor,
	}

	return server, nil
}