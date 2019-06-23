package main

import (
	"fmt"
	"net/http"
	"net/url"
	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
)

func handlerQRGenerator(c *gin.Context) {
	var png []byte

	var query GenerateCodeQuery
	
	if err := c.ShouldBind(&query); err != nil {
		response := GenericErrorResponse{
			Detail: "Missing Params",
			Code: "EMP",
		}
		response.Send(c)
		return
	}

	if !query.Validate() {
		response := GenericErrorResponse{
			Detail: "Invalid Input",
			Code: "EII",
		}
		response.Send(c)
		return
	}
	params := url.Values{}
	params.Add("secret", query.Secret)
	params.Add("issuer", ConstIssuer)

	payload := "otpauth://totp/TOTP?%s"
	payload = fmt.Sprintf(payload, params.Encode())

	png, err := qrcode.Encode(payload, qrcode.Medium, 256)

	if err != nil {
		response := GenericErrorResponse{
			Detail: "Internal Error",
			Code: "EGE",
		}
		response.Send(c)
		return	
	}
	c.Data(http.StatusOK, "image/png", png)
}

func handlerCodeGenerator(c *gin.Context) {
	var query GenerateCodeQuery
	
	if err := c.ShouldBind(&query); err != nil {
		response := GenericErrorResponse{
			Detail: "Missing Params",
			Code: "EMP",
		}
		response.Send(c)
		return
	}

	if !query.Validate() {
		response := GenericErrorResponse{
			Detail: "Invalid Input",
			Code: "EII",
		}
		response.Send(c)
		return
	}

	query.Secret = ConstSecretPrefix + query.Secret
	
	code, err := currentCode(query)
	
	if err != nil {
		response := GenericErrorResponse{
			Detail: "Invalid Secret",
			Code: "EIS",
		}
		response.Send(c)
		return
	}

	response := GetQueryResponse{
		Code: code,
	}
	
	response.Send(c)
}

func handleCodeVerification(c *gin.Context) {
	var query ValidateQuery;

	if err := c.ShouldBind(&query); err != nil {
		response := GenericErrorResponse{
			Detail: "Missing Params",
			Code: "EMP",
		}
		response.Send(c)
		return
	}
	if !query.Validate() {
		response := GenericErrorResponse{
			Detail: "Invalid Input",
			Code: "EII",
		}
		response.Send(c)
		return
	}

	query.Secret = ConstSecretPrefix + query.Secret
	result, err := verify(query)
	
	if err != nil {
		response := GenericErrorResponse{
			Detail: "Invalid Secret",
			Code: "EIS",
		}
		response.Send(c)
		return
	}

	response := ValidateCodeResponse{Valid: result}
	response.Send(c)
}


func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/code", handlerCodeGenerator)
	r.POST("/code", handlerCodeGenerator)

	r.GET("/qr", handlerQRGenerator)
	r.POST("/qr", handlerQRGenerator)

	r.GET("/verify", handleCodeVerification)
	r.POST("/verify", handleCodeVerification)
	return r
}

func serve(host string, port int) {
	r := setupRouter()
	r.Run(fmt.Sprintf("%s:%d", host, port)) // listen and serve on 0.0.0.0:8080
}