package routes

import (
	"payment/system/internal/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	// API_VERSION := "/api/v1"
	// route.POST(API_VERSION+"/users", controllers.Authenticate)
  auth := route.Group("/auth") {
    auth.POST("/login", controllers.Authenticate)
  }
}
