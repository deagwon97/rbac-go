package rest

import "github.com/gin-gonic/gin"

func AddRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/account")
	h, _ := NewHandler()
	router.POST("", h.AddUser)
	router.POST("/login", h.Login)
	router.POST("/valid", h.IsValid)
	router.POST("/renew", h.RenewToken)
	router.POST("/name/list", h.GetUserListName)
	// router.GET("/list", )
	// router.PATCH("/:id", )
	// router.DELETE("/:id", )
}
