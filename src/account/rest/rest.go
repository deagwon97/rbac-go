package rest

import "github.com/gin-gonic/gin"

func AddAccountRoutes(rg *gin.RouterGroup) {
	account := rg.Group("/account")
	h, _ := NewHandler()
	account.POST("", h.AddUser)
	account.POST("/login", h.Login)
	account.POST("/valid", h.IsValid)
	account.POST("/renew", h.RenewToken)
	// account.GET("/list", )
	// account.PATCH("/:id", )
	// account.DELETE("/:id", )
}
