package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "encoding/xml"
    "github.com/stretchr/testify/assert"
    "encoding/json"
    "net/url"
    "strings"
    "bytes"
)

func TestGetCodeQSJSON(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/code", nil)

    q := req.URL.Query()
    
    secret := "ABCDEF23ZWXYGHFM"

    q.Add("secret", secret)
    req.URL.RawQuery = q.Encode()

    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    response := GetQueryResponse{}

    assert.Equal(t, nil, json.Unmarshal(w.Body.Bytes(), &response))
    
    code, _ := currentCode(GenerateCodeQuery{Secret: ConstSecretPrefix + secret})
    assert.Equal(t, code, response.Code)
}

func TestGetCodeQSXML(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/code", nil)

    q := req.URL.Query()
    secret := "ABCDEF23ZWXYGHFM"
    q.Add("secret", secret)
    req.URL.RawQuery = q.Encode()

    req.Header.Add("Accept", "application/xml")
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    response := GetQueryResponse{}

    assert.Equal(t, nil, xml.Unmarshal(w.Body.Bytes(), &response))
    
    code, _ := currentCode(GenerateCodeQuery{Secret: ConstSecretPrefix + secret})
    assert.Equal(t, response.Code, code)
}


func TestGetCodeForm(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()

    secret := "ABCDEF23ZWXYGHFM"
    
    payload := url.Values {}
    payload.Set("secret", secret)

    req, _ := http.NewRequest("POST", "/code", strings.NewReader(payload.Encode()))
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    response := GetQueryResponse{}

    assert.Equal(t, nil, json.Unmarshal(w.Body.Bytes(), &response))
    
    code, _ := currentCode(GenerateCodeQuery{Secret: ConstSecretPrefix + secret})
    assert.Equal(t, code, response.Code)
}



func TestBadGetCodeForm(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()

    secret := "191919191919"
    
    payload := url.Values {}
    payload.Set("secret", secret)

    req, _ := http.NewRequest("POST", "/code", strings.NewReader(payload.Encode()))
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestBadGetCodeFormBadSecret(t *testing.T) {
    oldPf := ConstSecretPrefix

    router := setupRouter()

    w := httptest.NewRecorder()
    secret := "ABCDEFHHHDDZ"
    
    payload := url.Values {}
    payload.Set("secret", secret)

    req, _ := http.NewRequest("POST", "/code", strings.NewReader(payload.Encode()))
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    ConstSecretPrefix = "B1"
    router.ServeHTTP(w, req)
    ConstSecretPrefix = oldPf
    assert.Equal(t, http.StatusBadRequest, w.Code)
}



func TestBadGetCodeFormOverflow(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()

    secret := "ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789"
    
    payload := url.Values {}
    payload.Set("secret", secret)

    req, _ := http.NewRequest("POST", "/code", strings.NewReader(payload.Encode()))
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("Accept", "text/plain")

    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
    
    response := w.Body.String()

    assert.Equal(t, "E-EII: Invalid Input", response)
}


func TestBadGetCodeEmptySecret(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()

    secret := ""
    
    payload := url.Values {}
    payload.Set("secret", secret)

    req, _ := http.NewRequest("POST", "/code", strings.NewReader(payload.Encode()))
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("Accept", "text/plain")

    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
    
    response := w.Body.String()

    assert.Equal(t, "E-EMP: Missing Params", response)
}

func TestBadGetCodeEmpty(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()

    req, _ := http.NewRequest("POST", "/code", nil)
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("Accept", "text/plain")

    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusBadRequest, w.Code)
    
    response := w.Body.String()
    
    assert.Equal(t, response, "E-EMP: Missing Params")
}

func TestBadGetCodeEmptyGET(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()

    req, _ := http.NewRequest("GET", "/code", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
}


func TestGetCodeJSON(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()

    secret := "ABCDEF23ZWXYGHFM"
    
    payload := map[string]string{
        "secret" : secret,  
    }

    jsonValue, _ := json.Marshal(payload)

    req, _ := http.NewRequest("POST", "/code", bytes.NewBuffer(jsonValue))
    req.Header.Add("Content-Type", "application/json")

    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    response := GetQueryResponse{}

    assert.Equal(t, nil, json.Unmarshal(w.Body.Bytes(), &response))
    
    code, _ := currentCode(GenerateCodeQuery{Secret: ConstSecretPrefix + secret})
    assert.Equal(t, code, response.Code)
}

func TestGetCodeXML(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()

    secret := "ABCDEF23ZWXYGHFM"
    
    payload := GenerateCodeQuery{
        Secret: secret,  
    }

    xmlValue, _ := xml.MarshalIndent(payload, "", "")

    req, _ := http.NewRequest("POST", "/code", bytes.NewBuffer(xmlValue))
    req.Header.Add("Content-Type", "application/xml")

    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    response := GetQueryResponse{}

    assert.Equal(t, nil, json.Unmarshal(w.Body.Bytes(), &response))
    
    code, _ := currentCode(GenerateCodeQuery{Secret: ConstSecretPrefix + secret})
    assert.Equal(t, code, response.Code)
}


func estGetValidateCode(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()

    secret := ConstSecretPrefix + "A234567B"

    genQuery := GenerateCodeQuery{
        Secret: secret,
    }
    
    assert.Equal(t, true, genQuery.Validate())

    code, _ := currentCode(genQuery)
    
    payload := ValidateQuery{
        Secret: secret,
        Code: code,
    }

    jsonValue, _ := json.Marshal(payload)

    req, _ := http.NewRequest("POST", "/verify", bytes.NewBuffer(jsonValue))
    req.Header.Add("Content-Type", "application/json")

    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    response := ValidateCodeResponse{}

    assert.Equal(t, nil, json.Unmarshal(w.Body.Bytes(), &response))
    
    assert.Equal(t, true, response.Valid)
}
