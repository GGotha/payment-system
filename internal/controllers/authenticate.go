package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ta funcionando",
	})
}
