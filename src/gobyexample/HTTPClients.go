package main

import (
	"ad_opt_admin/src/advert/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	/*resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)

	//1创建文件
	f, _ := os.Create("/tmp/go.html")
	defer f.Close()
	for i := 0; scanner.Scan(); i++ {
		//fmt.Println(scanner.Text())
		//2写入文件
		f.WriteString(scanner.Text() + "\n")
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}*/

	client := &http.Client{}
	//读取Api数据
	url := "http://kq.ngking.com:8080/iclock/staff/transaction/?p=1&t=staff_transaction.html&UserID__id__exact=XXX&fromTime=&toTime="
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Cookie", "sessionid=5edb1f18c5a0cb334b42b2383c899e01")
	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}
