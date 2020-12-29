package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

type text struct {
    AlarmAmount int `json:"alarm_amount"`
    AlarmBotKey string `json:"alarm_bot_key"`
}

// 读取json文件
func main() {
    dir, _ := os.Getwd()
    fmt.Println("当前路径：", dir)
    f, err := ioutil.ReadFile("/Users/royod/go/src/golang/src/util/fileRed/text.json")
    if err != nil {
        fmt.Println("read fail", err)
    }
    var t text
    err = json.Unmarshal(f, &t)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(t.AlarmAmount)
    fmt.Println(t.AlarmBotKey)
    //fmt.Println(string(f))
    /*if !gjson.Valid(string(f)) {
    	fmt.Println(errors.New("invalid json"))
    }
    value := gjson.Get(string(f), "jieshuibang.merchant_no")
    println(value.String())
    m, ok := gjson.Parse(string(f)).Value().(map[string]interface{})
    if !ok {
    }
    var jes = m["jieshuibang"].(map[string]interface{})
    for k, i := range jes {
    	fmt.Println(k,i)
    }*/
}
