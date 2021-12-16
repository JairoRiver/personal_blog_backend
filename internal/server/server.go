package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func New() *Server {
	var server Server

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
