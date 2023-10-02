package api

import "github.com/gin-gonic/gin"

func (server *Server) setupRouter() {
	router := gin.Default()

	//autRoutes := router.Group("/")

	// User routes
	router.POST("/v1/user", server.createUser)
	router.GET("/v1/user/:id", server.getUser)
	router.GET("/v1/users", server.listUsers)
	router.PUT("/v1/user/:id", server.updateUser)
	router.DELETE("/v1/user/:id", server.deleteUser)

	server.router = router
}
