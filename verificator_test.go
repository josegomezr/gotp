package main
import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"fmt"
	"github.com/dgryski/dgoogauth"
)

func TestCurrentCode(t *testing.T){
	query := GenerateCodeQuery{
		Secret: ConstSecretPrefix + "A234567B",
	}
	
	assert.Equal(t, true, query.Validate())

	t0 := int64(time.Now().UTC().Unix() / 30)
	c := dgoogauth.ComputeCode(query.Secret, t0)

	code, _ := currentCode(query)
	assert.Equal(t, fmt.Sprintf("%06d", c), code)
}

func TestVerifyCode(t *testing.T){
	secret := ConstSecretPrefix + "A234567B"

	genQuery := GenerateCodeQuery{
		Secret: secret,
	}
	
	assert.Equal(t, true, genQuery.Validate())

	code, _ := currentCode(genQuery)
	
	valQuery := ValidateQuery{
		Secret: secret,
		Code: code,
	}

	assert.Equal(t, true, valQuery.Validate())

	result, err := verify(valQuery)

	assert.Equal(t, nil, err)
	assert.Equal(t, true, result)

	valQuery.Secret = valQuery.Secret + "1"
	
	assert.Equal(t, false, valQuery.Validate())
}
