package routes

import (
	"github.com/gin-gonic/gin"

	docs "rbac-go/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	accountRest "rbac-go/account/rest"
	rbacRest "rbac-go/rbac/rest"
)

func Run(address string) error {

	router := gin.New()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:8000",
		"http://localhost:8001",
		"http://localhost:3000",
		"http://localhost:3001",
		"https://rbac.deagwon.com"}
	router.Use(cors.New(config))

	router.Use(static.Serve("/admin", static.LocalFile("./admin/build", true)))

	docs.SwaggerInfo.BasePath = "/"

	v1 := router.Group("/")
	rbacRest.AddRoutes(v1)
	accountRest.AddRoutes(v1)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router.Run(address)
}
