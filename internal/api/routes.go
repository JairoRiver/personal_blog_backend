package api

import (
	"github.com/JairoRiver/personal_blog_backend/docs" // Swagger generated files
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func (server *Server) setupRouter() {
	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Personal Blog - API"
	docs.SwaggerInfo.Description = "Personal Blog - Post and Users API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/v1"
	//	@securityDefinitions.apiKey	JWT
	//	@in							header
	//	@name						token
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()

	//autRoutes := router.Group("/")
	apiRoutes := router.Group(docs.SwaggerInfo.BasePath)

	// User routes
	apiRoutes.POST("/user", server.createUser)
	apiRoutes.GET("/user/:id", server.getUser)
	apiRoutes.GET("/users", server.listUsers)
	apiRoutes.PUT("/user/:id", server.updateUser)
	apiRoutes.DELETE("/user/:id", server.deleteUser)

	// swagger
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	server.router = router
}
