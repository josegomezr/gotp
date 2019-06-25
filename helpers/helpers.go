package helpers
import (
	"strings"
    "regexp"
    "encoding/base32"
    "crypto/md5"
	"io"
)

// ValidBase32 ...
func ValidBase32(input string) (bool, error) {
    return regexp.MatchString("^(?:[A-Z2-7]{8})*(?:[A-Z2-7]{2}={6}|[A-Z2-7]{4}={4}|[A-Z2-7]{5}={3}|[A-Z2-7]{7}=)?$", input)
}

// FillSecret ...
func FillSecret(secret string) string {
    padding := len(secret) % 8
    
    if padding > 0 {
        secret = secret + strings.Repeat("=", 8 - padding)
    }

    secret = strings.ToUpper(secret)
    return secret
}

// OnlyOneStringOf ...
func OnlyOneStringOf(value1 string, value2 string) bool {
    if value1 == "" && value2 == "" {
        return false
    }
    
    if value1 != "" && value2 != "" {
        return false
    }

    return true
}


// SecretFromPayload ...
func SecretFromPayload(payload string) string {
	h := md5.New()
	io.WriteString(h, payload)
	return base32.StdEncoding.EncodeToString(h.Sum(nil))
}
