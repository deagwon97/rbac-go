package rest

import "github.com/gin-gonic/gin"

func AddRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/rbac")
	h, _ := NewHandler()
	router.GET("/role/list", h.GetRolesPage)
}
