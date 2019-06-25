package api
import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// CurrentCode_XXXXX
func TestCurrentCode_ALLGOOD(t *testing.T) {
	query := RequestGenerateCode{
		Secret: "ABCD23456ABCD234",
	}
	code, err := CurrentCode(query)
	
	assert.Nil(t,err)
	assert.NotEqual(t, "", code)
	assert.NotEqual(t, "-00001", code)
	assert.NotEqual(t, "000000", code)

	query.Secret = "1111111111"

	code, err = CurrentCode(query)
	
	assert.NotNil(t,err)
	assert.Equal(t, "", code)
}

// Verify_XXXXX
func TestVerify_ALLGOOD(t *testing.T) {
	secret := "ABCD23456ABCD234"
	genQuery := RequestGenerateCode{
		Secret: secret,
	}
	code, _ := CurrentCode(genQuery)
	
	query := RequestValidateCode{
		Secret: secret,
		Code: code,
	}

	valid, err := Verify(query)
	
	assert.Nil(t,err)
	assert.True(t, valid)
	
	query.Code = "00000A"

	valid, err = Verify(query)
	
	assert.NotNil(t,err)
	assert.False(t, valid)

	query.Code = "000000"
	query.Secret = "1111111"

	valid, err = Verify(query)
	
	assert.NotNil(t, err)
	assert.False(t, valid)
}


// QRPicture_XXXXX
func TestQRPicture(t *testing.T) {
	genQuery := RequestGenerateQR{
		Secret: "ABCD23456ABCD234",
	}
	
	png1, err := QRPicture(genQuery)
	
	assert.Nil(t,err)

	genQuery.Name = "Lorem Ipsum Dolor"
	
	png2, err := QRPicture(genQuery)
	
	assert.Nil(t, err)
	assert.NotEqual(t, png1, png2)
}
