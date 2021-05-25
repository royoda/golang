package main

import (
	"net/http"
	"os"
	"text/template"
)

/**
    *@Description TODO
    *@Date 2021/5/19 7:44 下午
**/

func main() {
	/*正则
	r, _ := regexp.Compile("^\\d{1,}$")
	res := r.MatchString("0012")
	fmt.Println(res)*/

	type Inventory struct {
		Material string
		Count    string
	}
	sweaters := Inventory{"wool", "a"}
	tmpl, err := template.New("test").Parse(`items are made of {{.Material}} {{.Count}} `)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
	//启动一个web服务器
	http.HandleFunc("/", handle)
	http.ListenAndServe("0.0.0.0:8000", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {

}
