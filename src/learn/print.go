package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"text/template"
	"time"
)

/**
    *@Description TODO
    *@Date 2021/5/10 2:37 下午
**/
func A() string {
	return "1"
}
func main() {
	fmt.Printf("%s\n", "\346\260\247\346\260\224\345\272\224\347\224\250")
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
	curTime := time.Now()
	fmt.Println(curTime)
	stTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2021-05-19 08:00:00", time.Local)
	fmt.Println()
	fmt.Println(stTime)
	fmt.Println(stTime.Unix())

	/*doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}*/
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
