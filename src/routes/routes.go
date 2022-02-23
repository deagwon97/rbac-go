package routes

import (
	"github.com/gin-gonic/gin"

	docs "rbac/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	accountRest "rbac/account/rest"
	contentRest "rbac/content/rest"
)

func Run(address string) error {

	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/"

	v1 := router.Group("/")
	contentRest.AddContentRoutes(v1)
	accountRest.AddAccountRoutes(v1)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router.Run(address)
}
