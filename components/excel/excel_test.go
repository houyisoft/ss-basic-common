package excel

import (
	"fmt"
	"testing"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func TestCreate(t *testing.T)  {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	err := f.SaveAs("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func TestRead(t *testing.T)  {
	f, err := excelize.OpenFile("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}


func TestAddChat(t *testing.T){
	f, err := excelize.OpenFile("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Insert a picture.
	err = f.AddPicture("Sheet1", "A2", "./image1.png", "")
	if err != nil {
		fmt.Println(err)
	}
	// Insert a picture to worksheet with scaling.
	err = f.AddPicture("Sheet1", "D2", "./image2.jpg", `{"x_scale": 0.5, "y_scale": 0.5}`)
	if err != nil {
		fmt.Println(err)
	}
	// Insert a picture offset in the cell with printing support.
	err = f.AddPicture("Sheet1", "H2", "./image3.gif", `{"x_offset": 15, "y_offset": 10, "print_obj": true, "lock_aspect_ratio": false, "locked": false}`)
	if err != nil {
		fmt.Println(err)
	}
	// Save the xlsx file with the origin path.
	err = f.Save()
	if err != nil {
		fmt.Println(err)
	}
}