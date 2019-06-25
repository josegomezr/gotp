package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/josegomezr/gotp/api"
)

// QRGenerator ...
func QRGenerator(c *gin.Context) {
	var png []byte

	var query api.RequestGenerateQR
	
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
	
	png, err := api.QRPicture(query)

	if err != nil {
		response := api.ResponseGenericError{
			Detail: "Internal Error",
			Code: "EGE",
			Error: err,
		}
		response.Send(c)
		return	
	}
	c.Data(http.StatusOK, "image/png", png)
}
