package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func handlerCodeGenerator(c *gin.Context) {
	var query GenerateCodeQuery;
	if err := c.ShouldBind(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing params"})
		return
	}
	if !query.validate() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Secret"})
		return
	}

	query.Secret = ConstSecretPrefix + query.Secret

	code, err := currentCode(query)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Secret"})
		return
	}
	
	acceptHeader := c.GetHeader("Accept")

	switch acceptHeader {
		case "application/xml":
			c.XML(http.StatusOK, GetQueryResponse{Code: code})
		break
		case "text/plain":
			c.String(http.StatusOK, code)
		break
		default:
			c.JSON(http.StatusOK, gin.H{"code": code,})
		break
	}
}

func outputNegotiation(c *gin.Context, statusCode int, payload interface{}) {
	acceptHeader := c.GetHeader("Accept")

	switch acceptHeader {
		case "application/xml":
			c.XML(statusCode, payload)
		break
		default:
			c.JSON(statusCode, payload)
		break
	}
}


func handleCodeVerification(c *gin.Context) {
	var query ValidateQuery;

	statusCode := http.StatusOK

	if err := c.ShouldBind(&query); err != nil {
		statusCode = http.StatusBadRequest
		outputNegotiation(c, statusCode, gin.H{"error": "missing params"})
		return
	}
	if !query.validate() {
		statusCode = http.StatusBadRequest
		outputNegotiation(c, statusCode, gin.H{"error": "Invalid Secret"})
		return
	}

	query.Secret = ConstSecretPrefix + query.Secret
	result, err := verify(query)
	
	if err != nil {
		statusCode = http.StatusBadRequest
		outputNegotiation(c, statusCode, gin.H{"error": "Invalid code"})
		return
	}
	
	outputNegotiation(c, statusCode, ValidateCodeResponse{Valid: result})
}


func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/code", handlerCodeGenerator)
	r.POST("/code", handlerCodeGenerator)
	r.GET("/verify", handleCodeVerification)
	r.POST("/verify", handleCodeVerification)
	return r
}

func serve(port int) {
	r := setupRouter()
	r.Run(fmt.Sprintf(":%d", port)) // listen and serve on 0.0.0.0:8080
}