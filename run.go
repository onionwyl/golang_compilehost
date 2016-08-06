package main

import (
    "fmt"
    "os/exec"
    "strings"
    "bytes"
    "log"
)

func main(){ 
    var out bytes.Buffer
    cmd := exec.Command("/bin/sh", "-c", "gcc a.c -o a.out &> err ")
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
    cmd = exec.Command("/bin/sh", "-c", "./a.out")
    cmd.Stdin = strings.NewReader("1")
    cmd.Stdout = &out
    cmd.Run()
    fmt.Printf("%s", out.String())
}
