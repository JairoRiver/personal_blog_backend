package api

import (
	"fmt"

	db "github.com/JairoRiver/personal_blog_backend/internal/db/sqlc"
	"github.com/JairoRiver/personal_blog_backend/pkg/token"
	"github.com/JairoRiver/personal_blog_backend/pkg/util"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP request for out bloging services
type Server struct {
	config     util.Config
	store      db.Querier
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store db.Querier) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return &server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
