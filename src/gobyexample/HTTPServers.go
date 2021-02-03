package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"math/rand"
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

//https://xuri.me/excelize/zh-hans/workbook.html#OpenReader
func excel(w http.ResponseWriter, req *http.Request) {
	file := excelize.NewFile()
	streamWriter, err := file.NewStreamWriter("Sheet1")
	if err != nil {
		fmt.Println(err)
	}
	styleID, err := file.NewStyle(`{"font":{"color":"#777777"}}`)
	if err != nil {
		fmt.Println(err)
	}
	if err := streamWriter.SetRow("A1", []interface{}{
		excelize.Cell{StyleID: styleID, Value: "Data"}}); err != nil {
		fmt.Println(err)
	}
	for rowID := 2; rowID <= 102400; rowID++ {
		row := make([]interface{}, 10)
		for colID := 0; colID < 10; colID++ {
			row[colID] = rand.Intn(640000)
		}
		cell, _ := excelize.CoordinatesToCellName(1, rowID)
		if err := streamWriter.SetRow(cell, row); err != nil {
			fmt.Println(err)
		}
	}
	if err := streamWriter.Flush(); err != nil {
		fmt.Println(err)
	}

	/*if err := file.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}*/

	w.Header().Set("Content-Disposition", "attachment; filename=Book1.xlsx")
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	if _, err := file.WriteTo(w); err != nil {
		fmt.Fprintf(w, err.Error())
	}
}
func main() {
	http.HandleFunc("/hello", hellohttp)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/excel", excel)

	http.ListenAndServe(":8090", nil)
}
