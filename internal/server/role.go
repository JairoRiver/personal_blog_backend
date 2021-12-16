package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) listRole(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "admin",
	})
}
