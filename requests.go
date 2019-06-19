package main
import (
    "gopkg.in/go-playground/validator.v9"
    "regexp"
    "strings"
)

type GenerateCodeQuery struct {
    Secret string `json:"secret" form:"secret" xml:"secret" validate:"required,alphanum,min=12,max=32" binding:"required"`
}

type ValidateQuery struct {
    Secret string `json:"secret" form:"secret" xml:"secret" validate:"required,alphanum,min=12,max=24" binding:"required"`
    Code  string `json:"code" form:"code" xml:"code"`
}

func validBase32(input string) bool {
    matched, error := regexp.MatchString("^([A-Z2-7=]+)$", input)
    return error == nil && matched
}

func (c *GenerateCodeQuery) validate() bool {
    v := validator.New()
    err := v.Struct(c)
    
    if err != nil {
        return false
    }
    
    padding := len(c.Secret) % 8
    if padding > 0 {
        c.Secret = c.Secret + strings.Repeat("=", padding)
    }
    c.Secret = strings.ToUpper(c.Secret)
    return validBase32(c.Secret)
}

func (c *ValidateQuery) validate() bool {
    v := validator.New()
    err := v.Struct(c)
    
    if err != nil {
        return false
    }
    padding := len(c.Secret) % 8
    if padding > 0 {
        c.Secret = c.Secret + strings.Repeat("=", padding)
    }
    c.Secret = strings.ToUpper(c.Secret)
    return validBase32(c.Secret)
}