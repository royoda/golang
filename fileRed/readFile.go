package main

import (
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
)

func main() {
	f, err := ioutil.ReadFile("text.json")
	if err != nil {
		fmt.Println("read fail", err)
	}
	//fmt.Println(string(f))
	if !gjson.Valid(string(f)) {
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
	}
}

