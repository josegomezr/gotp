package main

// GetQueryResponse response
type GetQueryResponse struct {
    Code  string `json:"code" form:"code" xml:"code"`
}


// ValidateCodeResponse response
type ValidateCodeResponse struct {
    Valid  bool `json:"valid" form:"valid" xml:"valid"`
}

