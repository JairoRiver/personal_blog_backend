package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	router *gin.Engine
}

func New() *Server {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hola Mundo",
		})
	})

	server := Server{
		router: router,
	}

	return &server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
