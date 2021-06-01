package common

import (
 	"fmt"
	"sort"
	"github.com/tealeg/xlsx"
)

const ROWS_LIMIT int = 1000

func sortReverse(slice map[string]int) []string {
	keys := make([]string, 0, len(slice))
	for key := range slice {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return slice[keys[i]] > slice[keys[j]] })

	return keys
}

func sumValues(values map[string]int) float64{
	var t int = 0
	for _, v := range values {
		t += v
	}
	return float64(t)
}

func percent(values map[string]int, v string) int64 {
	return int64(float64(values[v]) / sumValues(values) * 100)
}

func WriteGeral(f *xlsx.File, cell *xlsx.Cell, list map[string]int, sheetName string) {
	sh, _ := f.AddSheet("Geral")

	for index, key := range sortReverse(list) {
		cell, _ = sh.Cell(index, 0)
		cell.SetValue(key)
		cell, _ = sh.Cell(index, 1)
		cell.SetValue(list[key])
	}
}

func WriteRange(f *xlsx.File, sh *xlsx.Sheet, cell *xlsx.Cell, list map[string]map[string]int, sheetName string) {

	myStyle := xlsx.NewStyle()
	myStyle.Alignment.Horizontal = "right"
	myStyle.Fill.FgColor = "FFFFFF00"
	myStyle.Fill.PatternType = "solid"
	//myStyle.Font.Name = "Georgia"
	//myStyle.Font.Size = 11
	myStyle.Font.Bold = true
	myStyle.ApplyAlignment = true
	myStyle.ApplyFill = true
	myStyle.ApplyFont = true
	
	newColumn := xlsx.NewColForRange(2,2)
	newColumn.SetWidth(15)
	// we defined a style above, so let's assigm this style to all cells of the column
	//newColumn.SetStyle(myStyle)
	// now associate the sheet with this column

	sh, _ = f.AddSheet(sheetName)
	sh.SetColParameters(newColumn)

	index := 1 
	for cod, values := range list {
		cell, _ = sh.Cell(index, 0)
		cell.SetStyle(myStyle)
		cell.SetValue(cod)
		col := 2

		for _, v := range sortReverse(values) {
			cell, _ = sh.Cell(index, col)
			cell.SetValue(fmt.Sprintf("%d%s", percent(values, v), "%")) 
			cell, _ = sh.Cell(index, col+1)
			cell.SetValue(v)
			col += 2

			if col > ROWS_LIMIT { 
				break
			}
		}
		index++
	}
}