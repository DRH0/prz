package prz

import (
	"fmt"
	"math"
	"sort"
)

//func GroupByf64(table [][]float64, column)

func FilterByString(table [][]string, colnum int, id string) [][]string {
	var filtered [][]string
	for i := range table {
		if table[i][colnum] == id {
			filtered = append(filtered, table[i])
		}
	}
	return filtered
}

func GetColumnf64(table [][]float64, colnum int) []float64 {
	var column []float64
	for i := range table {
		column = append(column, table[i][colnum])
	}
	return column
}

func GetColumnString(table [][]string, colnum int) []string {
	var column []string
	for i := range table {
		column = append(column, table[i][colnum])
	}
	return column
}

// func ConvertColumnf64toString(column []float64) []string {
// 	var converted []string
// 	for i := range column {
// 		converted = append(converted, column[i])
// 	}
// }

func SumColumn(column []float64) float64 {
	var total float64 = 0
	for i := range column {
		total += column[i]
	}
	return total
}

func ChiSquared(actual []float64, guess []float64, filter float64) float64 {
	var fit float64 = 0
	for i := range actual {
		if actual[i] > filter {
			fit += (math.Pow(actual[i]-guess[i], 2) / guess[i])
		}
	}
	return fit
}

func Print2df64(table [][]float64) {
	for _, row := range table {
		for _, col := range row {
			fmt.Print(fmt.Sprintf("%.2f", col) + " ")
		}
		fmt.Print("\n")
	}
}

func Print2dString(table [][]string) {
	for _, row := range table {
		for _, col := range row {
			fmt.Print(col + " ")
		}
		fmt.Print("\n")
	}
}

func FwdFillString(table [][]string, testcolumn int, testvalue string, retrievecolumn int) {
	for i, row := range table {
		if row[testcolumn] == testvalue && i < 4 {
			table[i] = PrependString(row, row[retrievecolumn])
		} else if i != 0 {
			table[i] = PrependString(row, table[i-1][0])
		}
	}
}

func PrependString(row []string, s string) []string {
	row = append(row, "")
	copy(row[1:], row)
	row[0] = s
	return row
}

func Sort2dString(table [][]string, column int) {
	sort.SliceStable(table, func(i, j int) bool {
		return table[i][column] < table[j][column]
	})
}

func Sort2df64(table [][]float64, column int) {
	sort.SliceStable(table, func(i, j int) bool {
		return table[i][column] < table[j][column]
	})
}

//	convert string column to float
//	convert float column to string
//  filter by column equals string
//	filter by column comparison, perhaps a in-line function here (figure out)
//	merge on (vlookup style)
//	possibly sort on key and use sort.search (binary search)
//	average on (average ifs style)
