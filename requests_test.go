package main
import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidBase32(t *testing.T) {
	assert.Equal(t, true, validBase32("ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"))
	assert.Equal(t, false, validBase32("1"))
	assert.Equal(t, false, validBase32("9"))
}

func TestGenQuery(t *testing.T){
	query := GenerateCodeQuery{
		Secret: "ABCDEF23ZWXYGHFM",
	}
	assert.Equal(t, true, query.Validate())
	
	query.Secret = "ABC"
	assert.Equal(t, false, query.Validate())

	query.Secret = "1"
	assert.Equal(t, false, query.Validate())

}
func TestValQuery(t *testing.T){
	query := ValidateQuery{
		Secret: "ABCDEF23ZWXYGHFM",
	}
	assert.Equal(t, true, query.Validate())

	query.Secret = "1"
	assert.Equal(t, false, query.Validate())
	query.Secret = "ABC"
	assert.Equal(t, false, query.Validate())
}
