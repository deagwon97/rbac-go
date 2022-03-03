package error

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinError(c *gin.Context, err error) bool {
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return true
	}
	return false
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
