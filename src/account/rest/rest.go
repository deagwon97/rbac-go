package rest

import "github.com/gin-gonic/gin"

func AddContentRoutes(rg *gin.RouterGroup) {
	content := rg.Group("/content")
	h, _ := NewHandler()
	content.GET("/list", h.GetContents)

	content.GET("/:id", h.GetContent)
	content.POST("", h.AddContent)
	content.PATCH("/:id", h.UpdateContent)
	content.DELETE("/:id", h.DeleteContent)
}
