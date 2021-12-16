package server

import (
	"github.com/JairoRiver/personal_blog_backend/internal/util"
	db "github.com/JairoRiver/personal_blog_backend/pkg/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config util.Config
	router *gin.Engine
	store  *db.Queries
}

func New(config util.Config, store *db.Queries) *Server {
	server := Server{
		config: config,
		store:  store,
	}

	server.setupRouter()

	return &server
}

func (server *Server) setupRouter() {
	router := gin.Default()

	//db routes
	router.GET("/v1/role", server.listRole)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
