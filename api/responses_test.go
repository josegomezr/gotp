package api
import (
	"github.com/stretchr/testify/assert"
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"encoding/xml"
	"errors"
)

// ResponseGenericError_XXXXXXXX
func Test_ResponseGenericError_JSON(t *testing.T){
	response := ResponseGenericError{
		Detail: "Lorem ipsum dolor",
		Code: "E1234",
		Error: errors.New("Sample"),
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	c.Request, _ = http.NewRequest("GET", "http://example.com/", nil)
	c.Request.Header.Add("Accept", "application/json")
	response.Send(c)

	output := map[string]interface{}{}
	
	assert.Nil(t,json.Unmarshal(w.Body.Bytes(), &output))
	
	assert.Equal(t, output["code"], response.Code)
	assert.Equal(t, output["detail"], response.Detail)
	assert.NotContains(t, output, "error")
}

func Test_ResponseGenericError_PLAIN(t *testing.T){
	response := ResponseGenericError{
		Detail: "Lorem ipsum dolor",
		Code: "E1234",
		Error: errors.New("Sample"),
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	c.Request, _ = http.NewRequest("GET", "http://example.com/", nil)
	c.Request.Header.Add("Accept", "text/plain")
	response.Send(c)

	output := map[string]interface{}{}
	
	assert.Equal(t, fmt.Sprintf("E-%s: %s", response.Code, response.Detail), w.Body.String(), &output)
}

func Test_ResponseGenericError_XML1(t *testing.T){
	response := ResponseGenericError{
		Detail: "Lorem ipsum dolor",
		Code: "E1234",
		Error: errors.New("Sample"),
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	c.Request, _ = http.NewRequest("GET", "http://example.com/", nil)
	c.Request.Header.Add("Accept", "application/xml")
	response.Send(c)

	output := ResponseGenericError{}
	
	assert.Nil(t,xml.Unmarshal(w.Body.Bytes(), &output))
	
	assert.Equal(t, output.Code, response.Code)
	assert.Equal(t, output.Detail, response.Detail)
	assert.Equal(t, output.Error, nil)
	assert.NotEqual(t, output.Error, response.Error)
	// assert.NotContains(t, output, "error")
}

func Test_ResponseGenericError_XML2(t *testing.T){
	response := ResponseGenericError{
		Detail: "Lorem ipsum dolor",
		Code: "E1234",
		Error: errors.New("Sample"),
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	c.Request, _ = http.NewRequest("GET", "http://example.com/", nil)
	c.Request.Header.Add("Accept", "application/xml")
	response.Send(c)

	output := ResponseGenericError{}
	
	assert.Nil(t,xml.Unmarshal(w.Body.Bytes(), &output))
	
	assert.Equal(t, output.Code, response.Code)
	assert.Equal(t, output.Detail, response.Detail)
	assert.Equal(t, output.Error, nil)
	assert.NotEqual(t, output.Error, response.Error)
	// assert.NotContains(t, output, "error")
}

// ResponseCodeGeneration_XXXXXXXX
func Test_ResponseCodeGeneration_JSON(t *testing.T){
	response := ResponseCodeGeneration{
		Code: "000000",
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	c.Request, _ = http.NewRequest("GET", "http://example.com/", nil)
	c.Request.Header.Add("Accept", "application/json")
	response.Send(c)

	output := map[string]interface{}{}
	
	assert.Nil(t,json.Unmarshal(w.Body.Bytes(), &output))
	
	assert.Equal(t, output["code"], response.Code)
}

func Test_ResponseCodeGeneration_XML(t *testing.T){
	response := ResponseCodeGeneration{
		Code: "000000",
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	c.Request, _ = http.NewRequest("GET", "http://example.com/", nil)
	c.Request.Header.Add("Accept", "application/xml")
	response.Send(c)

	output := ResponseCodeGeneration{}
	
	assert.Nil(t,xml.Unmarshal(w.Body.Bytes(), &output))
	
	assert.Equal(t, response.Code, output.Code)
}

func Test_ResponseCodeGeneration_PLAIN(t *testing.T){
	response := ResponseCodeGeneration{
		Code: "000000",
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	c.Request, _ = http.NewRequest("GET", "http://example.com/", nil)
	c.Request.Header.Add("Accept", "text/plain")
	response.Send(c)

	assert.Equal(t, response.Code, w.Body.String())
}


// ResponseCodeVerification_XXXXXXXX
func Test_ResponseCodeVerification_JSON(t *testing.T){
	response := ResponseCodeVerification{
		Valid: true,
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	c.Request, _ = http.NewRequest("GET", "http://example.com/", nil)
	c.Request.Header.Add("Accept", "application/json")
	response.Send(c)

	output := map[string]interface{}{}
	
	assert.Nil(t,json.Unmarshal(w.Body.Bytes(), &output))
	
	assert.Equal(t, output["valid"], response.Valid)
}

func Test_ResponseCodeVerification_XML(t *testing.T){
	response := ResponseCodeVerification{
		Valid: true,
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	c.Request, _ = http.NewRequest("GET", "http://example.com/", nil)
	c.Request.Header.Add("Accept", "application/xml")
	response.Send(c)

	output := ResponseCodeVerification{}
	
	assert.Nil(t,xml.Unmarshal(w.Body.Bytes(), &output))
	
	assert.Equal(t, response.Valid, output.Valid)
}

func Test_ResponseCodeVerification_PLAIN_TRUE(t *testing.T){
	response := ResponseCodeVerification{
		Valid: true,
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	c.Request, _ = http.NewRequest("GET", "http://example.com/", nil)
	c.Request.Header.Add("Accept", "text/plain")
	response.Send(c)

	assert.Equal(t, "1", w.Body.String())
}

func Test_ResponseCodeVerification_PLAIN_FALSE(t *testing.T){
	response := ResponseCodeVerification{
		Valid: false,
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	c.Request, _ = http.NewRequest("GET", "http://example.com/", nil)
	c.Request.Header.Add("Accept", "text/plain")
	response.Send(c)

	assert.Equal(t, "-", w.Body.String())
}