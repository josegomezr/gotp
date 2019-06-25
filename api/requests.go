package api
import (
    "gopkg.in/go-playground/validator.v8"
    "github.com/josegomezr/gotp/helpers"
    "errors"
)

// RequestBase interface
type RequestBase struct {
    Secret string
    Payload string
}

// ErrorAmbiguousSecret ...
var ErrorAmbiguousSecret = errors.New("Ambiguous Secret Key Source")

// ErrorInvalidSecretFormat ...
var ErrorInvalidSecretFormat = errors.New("Invalid secret format")

// @todo add helpers.ValidBase32 validation on main

// RequestGenerateQR request
type RequestGenerateQR struct {
    Secret string `json:"secret" form:"secret" xml:"secret" binding:"omitempty,min=8,max=128"`
    Payload string `json:"payload" form:"payload" xml:"payload" binding:"omitempty,min=8,max=128"`
    Name string `json:"name" form:"name" xml:"name" binding:"omitempty,min=4,max=16"`
}

// RequestGenerateCode request
type RequestGenerateCode struct {
    Secret string `json:"secret" form:"secret" xml:"secret" binding:"omitempty,min=8,max=128"`
    Payload string `json:"payload" form:"payload" xml:"payload" binding:"omitempty,min=8,max=128"`
}

// RequestValidateCode request
type RequestValidateCode struct {
    Secret string `json:"secret" form:"secret" xml:"secret" binding:"omitempty,min=8,max=128"`
    Payload string `json:"payload" form:"payload" xml:"payload" binding:"omitempty,min=8,max=128"`
    Code  string `json:"code" form:"code" xml:"code" binding:"required,len=6,numeric"`
}

func (c *RequestGenerateQR) toBase() RequestBase {
    return RequestBase{
        Secret: c.Secret,
        Payload: c.Payload,
    }
}

func (c *RequestGenerateCode) toBase() RequestBase {
    return RequestBase{
        Secret: c.Secret,
        Payload: c.Payload,
    }
}

func (c *RequestValidateCode) toBase() RequestBase {
    return RequestBase{
        Secret: c.Secret,
        Payload: c.Payload,
    }
}

func validateSecretNPayload(c RequestBase) error {
    if !helpers.OnlyOneStringOf(c.Secret, c.Payload) {
        return ErrorAmbiguousSecret
    }

    c.Secret = helpers.FillSecret(c.Secret)
    
    matched, _ := helpers.ValidBase32(c.Secret)

    if ! matched {
        return ErrorInvalidSecretFormat
    }

    return nil
}

// Validate request validator
func (c *RequestGenerateQR) Validate() error {
    v := validator.New(&validator.Config{TagName: "binding"})
    err := v.Struct(c)
    
    if err != nil {
        return err
    }
        
    err = validateSecretNPayload(c.toBase())
    
    if err != nil {
        return err
    }
    
    if c.Secret == "" {
        c.Secret = helpers.SecretFromPayload(c.Payload)
    }

    return nil
}

// Validate request validator
func (c *RequestGenerateCode) Validate() error {
    v := validator.New(&validator.Config{TagName: "binding"})
    err := v.Struct(c)
    
    if err != nil {
        return err
    }

    err = validateSecretNPayload(c.toBase())
    
    if err != nil {
        return err
    }
    
    if c.Secret == "" {
        c.Secret = helpers.SecretFromPayload(c.Payload)
    }

    return nil
}

// Validate request validator
func (c *RequestValidateCode) Validate() error {
    v := validator.New(&validator.Config{TagName: "binding"})
    err := v.Struct(c)
    
    if err != nil {
        return err
    }
    
    err = validateSecretNPayload(c.toBase())
    
    if err != nil {
        return err
    }
    
    if c.Secret == "" {
        c.Secret = helpers.SecretFromPayload(c.Payload)
    }

    return nil
}
