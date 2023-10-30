package api

import (
	"log"

	"github.com/JairoRiver/personal_blog_backend/docs" // Swagger generated files
	"github.com/JairoRiver/personal_blog_backend/pkg/util"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func (server *Server) setupRouter() {
	config, err := util.LoadConfig(".", "app")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Personal Blog - API"
	docs.SwaggerInfo.Description = "Personal Blog - Post and Users API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = config.HostName
	docs.SwaggerInfo.BasePath = "/v1"
	//	@securityDefinitions.apiKey	JWT
	//	@in							header
	//	@name						authorization
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	//autRoutes := router.Group("/")
	apiRoutes := router.Group(docs.SwaggerInfo.BasePath)
	authRoutes := apiRoutes.Group("").Use(authMiddleware(server.tokenMaker))

	// User routes
	authRoutes.POST("/user", server.createUser)
	authRoutes.GET("/user/:id", server.getUser)
	authRoutes.GET("/users", server.listUsers)
	authRoutes.PUT("/user/:id", server.updateUser)
	authRoutes.DELETE("/user/:id", server.deleteUser)
	apiRoutes.POST("/login", server.loginUser)

	// Categories routes
	authRoutes.POST("/category", server.createCategory)
	apiRoutes.GET("/category/:id", server.getCategory)
	apiRoutes.GET("/categories", server.listCategories)
	authRoutes.PUT("/category/:id", server.updateCategory)
	authRoutes.DELETE("/category/:id", server.deleteCategory)

	// Tags routes
	authRoutes.POST("/tag", server.createTag)
	apiRoutes.GET("/tag/:id", server.getTag)
	apiRoutes.GET("/tags", server.listTags)
	authRoutes.DELETE("/tag/:id", server.deleteTag)

	// Post routes admin
	authRoutes.POST("/admin/post", server.createPost)
	authRoutes.GET("/admin/post/:id", server.getPostByIdPrivate)
	authRoutes.GET("/admin/category-post/:id", server.getPostByCategoryPrivate)
	authRoutes.GET("/admin/tag-post/:id", server.getPostByTagPrivate)
	authRoutes.GET("/admin/posts", server.listPostsPrivate)
	authRoutes.PUT("/admin/post/{id}", server.updatePost)
	authRoutes.DELETE("/admin/post/{id}", server.deletePost)

	// Post routes public
	apiRoutes.GET("/post/:id", server.getPostByIdPublic)
	apiRoutes.GET("/category-post/:id", server.getPostByCategoryPublic)
	apiRoutes.GET("tag-post/:id", server.getPostByTagPublic)
	apiRoutes.GET("/posts", server.listPostsPublic)

	// PostTag routes
	authRoutes.POST("/admin/post-tag", server.createPostTag)
	apiRoutes.DELETE("/admin/post-tag/:id", server.deletePostTag)

	// swagger
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	server.router = router
}
