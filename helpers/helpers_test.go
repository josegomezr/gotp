package helpers
import (
	"github.com/stretchr/testify/assert"
	"testing"
	"strings"
)

// SecretFromPayload_XXXXX
func TestSecretFromPayload(t *testing.T) {
	secret := "LOREM IPSUM DOLOR"
	assert.Equal(t, "5747YKWZODSK4UDV7FGBHJRLYY======", SecretFromPayload(secret))
}

// OnlyOneStringOf_XXXXX
func TestOnlyOneStringOf(t *testing.T) {
	// only one
	assert.True(t, OnlyOneStringOf("", "ABC"))
	assert.True(t, OnlyOneStringOf("ABC", ""))
	
	// both
	assert.False(t, OnlyOneStringOf("", ""))
	assert.False(t, OnlyOneStringOf("A", "B"))
}

// ValidBase32_XXXXX
func TestValidBase32(t *testing.T) {
	// full alphabet
	matched, _ := ValidBase32("ABCDEFGHIJKLMNOQRSTUVWXYZ234567=")
	assert.True(t, matched)
	
	// not correctly balanced
	matched, _ = ValidBase32("ABCDEFGHIJKLMNOQRSTUVWXYZ234567")
	assert.False(t, matched)
	
	// bad alphabet
	matched, _ = ValidBase32("ABCDEFG1")
	assert.False(t, matched)

	// One pad
	matched, _ = ValidBase32("ABCDEFG=")
	assert.True(t, matched)

	// Two pads
	matched, _ = ValidBase32("ABCDEF==")
	assert.False(t, matched)

	// Three pads
	matched, _ = ValidBase32("ABCDE===")
	assert.True(t, matched)

	// Four pads
	matched, _ = ValidBase32("ABCD====")
	assert.True(t, matched)
	
	// Five pads
	matched, _ = ValidBase32("ABC=====")
	assert.False(t, matched)

	// Six pads
	matched, _ = ValidBase32("AB======")
	assert.True(t, matched)

	// Seven pads
	matched, _ = ValidBase32("A=======")
	assert.False(t, matched)

	// 8 pads
	matched, _ = ValidBase32("========")
	assert.False(t, matched)
	
	// 9 pads
	matched, _ = ValidBase32("ABCDEFG=========")
	assert.False(t, matched)
}

// FillSecret_XXXXX
func TestFillSecret(t *testing.T) {
	secret := "ABCDEFGH"
	assert.Equal(t, FillSecret(secret), secret)
	
	secret = "ABCDEFGHA"
	assert.NotEqual(t, FillSecret(secret), secret)
	secret = "ABCDEFGH"
	for i := 0; i < 8; i++ {
		input := secret + strings.Repeat("A", i+1)
		expected := input + strings.Repeat("=", 7-i)

		assert.Equal(t, FillSecret(input), expected)
	}
}