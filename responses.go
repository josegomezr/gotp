package main
import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

// GenericErrorResponse response
type GenericErrorResponse struct {
	Detail string `json:"detail" form:"detail" xml:"detail"`
	Code string `json:"code" form:"code" xml:"code"`
}

// Send response rendering
func (r *GenericErrorResponse) Send(c *gin.Context) {
	acceptHeader := c.GetHeader("Accept")

	switch acceptHeader {
		case "application/xml":
			c.XML(http.StatusBadRequest, r)
		break
        case "text/plain":
            c.String(http.StatusBadRequest, fmt.Sprintf("E-%s: %s", r.Code, r.Detail))
		break
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"code": r.Code,
				"detail": r.Detail,
			})
		break
	}
}

// GetQueryResponse response
type GetQueryResponse struct {
    Code  string `json:"code" form:"code" xml:"code"`
}

// Send response rendering
func (r *GetQueryResponse) Send(c *gin.Context) {
    acceptHeader := c.GetHeader("Accept")

	switch acceptHeader {
		case "application/xml":
			c.XML(http.StatusOK, r)
		break
        case "text/plain":
            c.String(http.StatusOK, r.Code)
		break
		default:
			c.JSON(http.StatusOK, gin.H{"code": r.Code})
		break
	}
}

// ValidateCodeResponse response
type ValidateCodeResponse struct {
    Valid  bool `json:"valid" form:"valid" xml:"valid"`
}

// Send response rendering
func (r *ValidateCodeResponse) Send(c *gin.Context) {
    acceptHeader := c.GetHeader("Accept")

	switch acceptHeader {
		case "application/xml":
			c.XML(http.StatusOK, r)
		break
        case "text/plain":
            if r.Valid {
                c.String(http.StatusOK, "1")
            } else {
                c.String(http.StatusOK, "-")
            }
		break
		default:
			c.JSON(http.StatusOK, gin.H{"valid": r.Valid})
		break
	}
}
