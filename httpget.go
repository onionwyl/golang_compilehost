package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type submission_s struct{
	Sid int
	Uid int
	Submit_time string
	Code string
	Lang string
	Input string
	Output string
	Err_info string
}

func httpGet() {
	var content submission_s
    resp, err := http.Get("http://localhost:8080/api/getcode")
    if err != nil {
        // handle error
    }
 
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }
 	err = json.Unmarshal(body, &content)
    fmt.Println(content.Code);
}

func main() {
    httpGet()
}