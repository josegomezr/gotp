package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/josegomezr/gotp/api"
)

// CodeVerification ...
func CodeVerification(c *gin.Context) {
	var query api.RequestValidateCode;

	if err := c.ShouldBind(&query); err != nil {
		response := api.ResponseGenericError{
			Detail: "Missing Params",
			Code: "EMP",
			Error: err,
		}
		response.Send(c)
		return
	}
	
	if err := query.Validate(); err != nil {
		response := api.ResponseGenericError{
			Detail: "Invalid Input",
			Code: "EII",
			Error: err,
		}
		response.Send(c)
		return
	}

	result, err := api.Verify(query)
	
	if err != nil {
		details := "Invalid Secret"
		code := "EIS"
		
		if err == api.ErrorExpiredCode {
			details = "Expired Code"
			code = "EEC"
		}

		response := api.ResponseGenericError{
			Detail: details,
			Code: code,
			Error: err,
		}

		fmt.Println(err)

		response.Send(c)
		return
	}

	response := api.ResponseCodeVerification{Valid: result}
	response.Send(c)
}