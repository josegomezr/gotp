package api
import (
	"fmt"
	"time"
	"errors"
	"github.com/dgryski/dgoogauth"
	"net/url"
	qrcode "github.com/skip2/go-qrcode"
)

// ErrorExpiredCode ...
var ErrorExpiredCode = errors.New("Invalid/Expired code")

// ErrorInvalidSecret ...
var ErrorInvalidSecret = errors.New("Invalid Secret Key")


// CurrentCode ...
func CurrentCode(query RequestGenerateCode) (string, error) {
	t0 := int64(time.Now().UTC().Unix() / 30)
	c := dgoogauth.ComputeCode(query.Secret, t0)

	if c == -1 {
		return "", ErrorInvalidSecret
	}

	return fmt.Sprintf("%06d", c), nil
}

// Verify ...
func Verify(query RequestValidateCode) (bool, error) {
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

	if result == false {
		return false, ErrorExpiredCode
	}

	return true, nil
}


// QRPicture ...
func QRPicture(query RequestGenerateQR) ([]byte, error) {
	if query.Name == "" {
		query.Name = "GOTP"
	}
	params := url.Values{}
	secret := query.Secret
	
	params.Add("secret", secret)
	params.Add("issuer", query.Name)
	payload := "otpauth://totp/TOTP?%s"
	payload = fmt.Sprintf(payload, params.Encode())
	// fmt.Println(payload)
	return qrcode.Encode(payload, qrcode.Medium, 256)
}
