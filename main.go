package main

import (
    "bytes"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "regexp"
    "log"
)

func main() {
    imgPath := "https://www.baidu.com/img/bd_logo1.png"
    reg, _ := regexp.Compile(`(\w|\d|_)*.png`)
    name := reg.FindStringSubmatch(imgPath)[0]
    fmt.Println(name)
    resp, err := http.Get(imgPath)
    handleErr(err)
    body, err := ioutil.ReadAll(resp.Body)
    handleErr(err)
    out, err := os.Create(name)
    handleErr(err)
    io.Copy(out, bytes.NewReader(body))
}

func handleErr(err error) {
    if err != nil {
        log.Println(err)
        return
    }
}