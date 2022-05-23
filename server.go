package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/josegomezr/gotp/handlers"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/code", handlers.CodeGenerator)
	r.POST("/code", handlers.CodeGenerator)

	r.GET("/qr", handlers.QRGenerator)
	r.POST("/qr", handlers.QRGenerator)

	r.GET("/verify", handlers.CodeVerification)
	r.POST("/verify", handlers.CodeVerification)
	return r
}

func serve(host string, port int) {
	r := setupRouter()
	fmt.Printf("Starting GOTP HTTP Server listening at: %s:%d\n", host, port)
	r.Run(fmt.Sprintf("%s:%d", host, port)) // listen and serve on 0.0.0.0:8080
}
