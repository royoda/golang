package main

import (
	"fmt"
	"net/http"
	"time"
)

func hellocontext(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")

	case <-ctx.Done(): //取消连接，终端取消使用ctrl+c

		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)

	}
}

func main() {

	http.HandleFunc("/hello", hellocontext)
	http.ListenAndServe(":8090", nil)
}
