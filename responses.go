package main

type GetQueryResponse struct {
    Code  string `json:"code" form:"code" xml:"code"`
}


type ValidateCodeResponse struct {
    Valid  bool `json:"valid" form:"valid" xml:"valid"`
}

