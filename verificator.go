package main
import (
	"fmt"
	"time"
	"errors"
	"regexp"
	"github.com/dgryski/dgoogauth"
)


func validBase32(input string) bool {
    length := len(input)

    if (length % 8) > 4 {
    	return false
	}
	
    matched, err := regexp.MatchString("^(?:[A-Z2-7]{8})*(?:[A-Z2-7]{2}={6}|[A-Z2-7]{4}={4}|[A-Z2-7]{5}={3}|[A-Z2-7]{7}=)?$", input)
	return err == nil && matched
}

func currentCode(query GenerateCodeQuery) (string, error) {
	t0 := int64(time.Now().UTC().Unix() / 30)
	c := dgoogauth.ComputeCode(query.Secret, t0)

	if c == -1 {
		return "", errors.New("Invalid Secret")
	}

	return fmt.Sprintf("%06d", c), nil
}

func verify(query ValidateQuery) (bool, error) {
	authconf := dgoogauth.OTPConfig{
		Secret:       query.Secret,
		WindowSize:   2,
		HotpCounter: 0,
		UTC: true,
	}
	result, err := authconf.Authenticate(query.Code)
	
	if err != nil {
		return false, err
	}

	return result, nil
}