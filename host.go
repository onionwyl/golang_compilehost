package main

import (
    "fmt"
    "os/exec"
    "strings"
    "bytes"
    "log"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "time"
    "os"
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

func httpGet() submission_s {
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
    //fmt.Println(content.Code);
    return content
}

func main() {
    for{
        submission := httpGet()
        fmt.Println(submission.Code)
        os.Remove("tmp.c")
        os.Remove("tmp.out")
        os.Remove("err")
        fout, _ := os.Create("tmp.c")
        defer fout.Close()
        fout.WriteString(submission.Code)
        var out bytes.Buffer
        cmd := exec.Command("/bin/sh", "-c", "gcc tmp.c -o tmp.out &> err ")
        cmd.Stdout = &out
        err := cmd.Run()
        if err != nil {
            log.Fatal(err)
        }
        cmd = exec.Command("/bin/sh", "-c", "./tmp.out")
        cmd.Stdin = strings.NewReader(submission.Input)
        cmd.Stdout = &out
        cmd.Run()
        fmt.Printf("%s", out.String())
        fin, _ := os.Open("err")
        defer fin.Close()
        fd, _ := ioutil.ReadAll(fin)
        fmt.Printf("%s", fd)
        time.Sleep(time.Second)
    }
}