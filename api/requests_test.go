package api
import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// RequestGenerateQR_XXXXXXXX
func Test_RequestGenerateQR_WithSecret(t *testing.T){
	query := RequestGenerateQR{
		Secret: "ABCDEF23ZWXYGHFM",
	}
	assert.Nil(t,query.Validate())
	query.Secret = "ABC"
	assert.NotNil(t,query.Validate())

	query.Secret = "1"
	assert.NotNil(t,query.Validate())
	
	query.Secret = "111111111111"
	assert.NotNil(t,query.Validate())

	query = RequestGenerateQR{
		Payload: "Lorem ipsum dolor",
	}

	assert.Nil(t,query.Validate())
	query.Payload = "ABC"
	assert.NotNil(t,query.Validate())

	query = RequestGenerateQR{
		Secret: "ABCDEF23ZWXYGHFM",
		Payload: "Lorem ipsum dolor",
	}
	assert.NotNil(t,query.Validate())
}


// RequestGenerateCode_XXXXXXXX
func Test_RequestGenerateCode_WithSecret(t *testing.T){
	query := RequestGenerateCode{
		Secret: "ABCDEF23ZWXYGHFM",
	}
	assert.Nil(t,query.Validate())
	query.Secret = "ABC"
	assert.NotNil(t,query.Validate())

	query.Secret = "1"
	assert.NotNil(t,query.Validate())
	
	query.Secret = "111111111111"
	assert.NotNil(t,query.Validate())

	query = RequestGenerateCode{
		Payload: "Lorem ipsum dolor",
	}

	assert.Nil(t,query.Validate())
	query.Payload = "ABC"
	assert.NotNil(t,query.Validate())

	query = RequestGenerateCode{
		Secret: "ABCDEF23ZWXYGHFM",
		Payload: "Lorem ipsum dolor",
	}
	assert.NotNil(t,query.Validate())
}


// RequestValidateCode_XXXXXXXX
func Test_RequestValidateCode_WithSecret(t *testing.T){
	query := RequestValidateCode{
		Code: "000000",
		Secret: "ABCDEF23ZWXYGHFM",
	}
	assert.Nil(t,query.Validate())
	query.Secret = "ABC"
	assert.NotNil(t,query.Validate())

	query.Secret = "1"
	assert.NotNil(t,query.Validate())
	
	query.Secret = "111111111111"
	assert.NotNil(t,query.Validate())

	query = RequestValidateCode{
		Code: "000000",
		Payload: "Lorem ipsum dolor",
	}

	assert.Nil(t,query.Validate())
	query.Payload = "ABC"
	assert.NotNil(t,query.Validate())

	query = RequestValidateCode{
		Secret: "ABCDEF23ZWXYGHFM",
		Payload: "Lorem ipsum dolor",
	}
	assert.NotNil(t,query.Validate())
	
	query = RequestValidateCode{
		Code: "00000A",
		Secret: "ABCDEF23ZWXYGHFM",
	}
	assert.NotNil(t,query.Validate())
}
