package main
// import (
// 	"github.com/stretchr/testify/assert"
// 	"testing"
// 	"time"
// 	"fmt"
// 	"github.com/dgryski/dgoogauth"
// )

// func TestCurrentCode(t *testing.T){
// 	query := RequestGenerateCode{
// 		Secret: ConstSecretPrefix + "A234567B",
// 	}
	
// 	assert.True(t, query.Validate())

// 	t0 := int64(time.Now().UTC().Unix() / 30)
// 	c := dgoogauth.ComputeCode(query.Secret, t0)

// 	code, _ := currentCode(query)
// 	assert.Equal(t, fmt.Sprintf("%06d", c), code)
// }

// func TestVerifyCode(t *testing.T){
// 	secret := ConstSecretPrefix + "A234567B"

// 	genQuery := RequestGenerateCode{
// 		Secret: secret,
// 	}
	
// 	assert.True(t, genQuery.Validate())

// 	code, _ := currentCode(genQuery)
	
// 	valQuery := RequestValidateCode{
// 		Secret: secret,
// 		Code: code,
// 	}

// 	assert.True(t, valQuery.Validate())

// 	result, err := verify(valQuery)

// 	assert.Nil(t,err)
// 	assert.True(t, result)

// 	valQuery.Secret = valQuery.Secret + "1"
	
// 	assert.False(t, valQuery.Validate())
// }
