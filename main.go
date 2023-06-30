package main

import (
	"log"

	// "github.com/ggotha/payment-system/core/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Default()

	startGin()
}

func startGin() {
	router := gin.New()

	API_PORT := ":8080"

	routes.Routes(router)
	// routes.Authenticate(router)
	// routes.routes(router)
	// r.POST(API_VERSION+"/users", controllers.Authenticate)

	router.Run(API_PORT)
}
