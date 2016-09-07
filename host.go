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
    Sid string 
    Submit_time string
    Code string
    Lang string
    Input string
    Output string
    Err_info string
}

func httpGet() submission_s {
    var content submission_s
    resp, err := http.Get("http://www.onionwyl.cn/api/getcode")
    if err != nil {
        // handle error
    }
 
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }
    //fmt.Printf("%s\n", body)
    err = json.Unmarshal(body, &content)
    //fmt.Println(content.Sid);
    return content
}

func main() {
    for{
        submission := httpGet()
        if submission.Sid == ""{
            time.Sleep(time.Second)
            continue
        }
        fmt.Println(submission)
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
           // log.Fatal(err)
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
        submission.Output = out.String()
        submission.Err_info = string(fd)
        time.Sleep(time.Second)
        j, _ := json.Marshal(submission)
        body := bytes.NewBuffer([]byte(j))
        resp, err := http.Post("http://www.onionwyl.cn/api/putresult", "application/json;charset=utf-8", body)
        if err != nil {
            log.Fatal(err)
            return
        }
       // result, _ := ioutil.ReadAll(resp.Body)
        defer resp.Body.Close()
       // fmt.Printf("%s", result)
    }
}
