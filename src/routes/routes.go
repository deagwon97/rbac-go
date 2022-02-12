package routes

import (
	"github.com/gin-gonic/gin"

	docs "rbac/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	content_rest "rbac/content/rest"
)

func Run(address string) error {

	router := gin.Default()

	docs.SwaggerInfo_swagger.BasePath = "/"

	v1 := router.Group("/")
	content_rest.AddContentRoutes(v1)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router.Run(address)
}
