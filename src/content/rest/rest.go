package rest

import "github.com/gin-gonic/gin"

func AddRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/content")
	h, _ := NewHandler()
	router.GET("/list", h.GetContents)

	router.GET("/:id", h.GetContent)
	router.POST("", h.AddContent)
	router.PATCH("/:id", h.UpdateContent)
	router.DELETE("/:id", h.DeleteContent)
}
