package main
import (
    "gopkg.in/go-playground/validator.v9"
    "strings"
)

// GenerateCodeQuery request
type GenerateCodeQuery struct {
    Secret string `json:"secret" form:"secret" xml:"secret" validate:"required,alphanum,min=8,max=32" binding:"required"`
}

// ValidateQuery request
type ValidateQuery struct {
    Secret string `json:"secret" form:"secret" xml:"secret" validate:"required,alphanum,min=8,max=32" binding:"required"`
    Code  string `json:"code" form:"code" xml:"code"`
}

func fillSecret(secret string) string {
    padding := len(secret) % 8
    
    if padding > 0 {
        secret = secret + strings.Repeat("=", 8 - padding)
    }
    secret = strings.ToUpper(secret)
    return secret
}

// Validate request validator
func (c *GenerateCodeQuery) Validate() bool {
    v := validator.New()
    err := v.Struct(c)
    
    if err != nil {
        return false
    }
    
    c.Secret = fillSecret(c.Secret)
    
    return validBase32(c.Secret)
}

// Validate request validator
func (c *ValidateQuery) Validate() bool {
    v := validator.New()
    err := v.Struct(c)
    
    if err != nil {
        return false
    }
    
    c.Secret = fillSecret(c.Secret)

    return validBase32(c.Secret)
}
