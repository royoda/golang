package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func hellohttp(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n") //将hello 写入response对象中
}

func headers(w http.ResponseWriter, req *http.Request) {
	fmt.Println(strconv.Atoi(req.URL.Query().Get("page")))
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hellohttp)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
