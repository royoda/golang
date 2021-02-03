package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"math/rand"
)

func main() {
	/*header := []interface{}{"Small", "Normal", "Large", "Apple", "Orange", "Pear"}
	dir, _ := os.Getwd()
	fmt.Println(dir)
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")
	f.SetSheetRow("Sheet1", "A1", &header)
	var indexNum = 1
	for i := 1; i< 10; i++ {
		indexNum = i + 1
		local := fmt.Sprintf("A%d",indexNum)
		f.SetSheetRow("Sheet1", local, &[]interface{}{1, 2, 3, 4, 5, "sadas"})
	}
	local := fmt.Sprintf("A%d",indexNum + 1)
	f.SetCellFormula("Sheet1", local, "=SUM(A2:A9)")
	// Set value of a cell.
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs(dir + "/Book1.xlsx"); err != nil {
		fmt.Println(err)
	}*/

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
	if err := file.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
