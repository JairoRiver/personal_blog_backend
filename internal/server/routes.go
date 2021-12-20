package server

import "github.com/gin-gonic/gin"

func (server *Server) setupRouter() {
	router := gin.Default()

	//role routes
	router.POST("/v1/role", server.createRole)
	router.GET("/v1/role/:id", server.getRole)
	router.GET("/v1/roles", server.listRoles)
	router.PUT("/v1/role/:id", server.updateRole)
	router.DELETE("/v1/role/:id", server.deleteRole)

	//user routes
	router.POST("/v1/user", server.createUser)
	router.GET("/v1/user/:id", server.getUser)
	router.GET("/v1/users", server.listUsers)
	router.PUT("/v1/user/:id", server.updateUser)
	router.DELETE("/v1/user/:id", server.deleteUser)

	server.router = router
}
