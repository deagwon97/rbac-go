package routes

import (
	"github.com/gin-gonic/gin"

	docs "rbac-go/docs"

	"github.com/gin-gonic/contrib/static"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	accountRest "rbac-go/account/rest"
	rbacRest "rbac-go/rbac/rest"
)

func Run(address string) error {

	router := gin.New()

	router.Use(CORSMiddleware())
	router.Use(static.Serve("/admin", static.LocalFile("./admin/build", true)))

	docs.SwaggerInfo.BasePath = "/"

	v1 := router.Group("/")
	rbacRest.AddRoutes(v1)
	accountRest.AddRoutes(v1)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router.Run(address)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
