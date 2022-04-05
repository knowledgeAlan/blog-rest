package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func DealHttpResponse(c *gin.Context, data gin.H) {
	c.JSON(http.StatusOK, data)
}