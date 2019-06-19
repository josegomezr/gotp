package main
import (
	"fmt"
	"time"
	"errors"
	"github.com/dgryski/dgoogauth"
)

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