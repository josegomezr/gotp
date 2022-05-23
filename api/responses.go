package api
import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"encoding/json"
)

// ResponseGenericError response
type ResponseGenericError struct {
	Detail string `json:"detail" form:"detail" xml:"detail"`
	Code string `json:"code" form:"code" xml:"code"`
	Error error
}

// ResponseCodeGeneration response
type ResponseCodeGeneration struct {
    Code  string `json:"code" form:"code" xml:"code"`
}

// ResponseCodeVerification response
type ResponseCodeVerification struct {
    Valid  bool `json:"valid" form:"valid" xml:"valid"`
}

// Send response rendering
func (r *ResponseGenericError) Send(c *gin.Context) {
	acceptHeader := c.GetHeader("Accept")

	switch acceptHeader {
		case "text/xml":
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

// Send response rendering
func (r *ResponseCodeGeneration) Send(c *gin.Context) {
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

func (r *ResponseCodeGeneration) SendText(format string) string {
	switch format {
		case "json":
			res, _ := json.Marshal(r)
			return string(res)
		break
        case "simple":
            return r.Code
		break
	}
	return fmt.Sprintf("Code: %s", r.Code)
}

// Send response rendering
func (r *ResponseCodeVerification) Send(c *gin.Context) {
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
