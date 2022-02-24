package routes

import (
	"github.com/gin-gonic/gin"

	docs "rbac-go/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	accountRest "rbac-go/account/rest"
	contentRest "rbac-go/content/rest"
	rbacRest "rbac-go/rbac/rest"
)

func Run(address string) error {

	router := gin.New()

	docs.SwaggerInfo.BasePath = "/"

	v1 := router.Group("/")
	rbacRest.AddRoutes(v1)
	contentRest.AddRoutes(v1)
	accountRest.AddRoutes(v1)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router.Run(address)
}
