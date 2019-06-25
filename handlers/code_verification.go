package handlers

import (
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
		response := api.ResponseGenericError{
			Detail: "Invalid Secret",
			Code: "EIS",
			Error: err,
		}
		response.Send(c)
		return
	}

	response := api.ResponseCodeVerification{Valid: result}
	response.Send(c)
}