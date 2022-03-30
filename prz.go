package prz

import (
	"math"
)

func GetColumnf64(table [][]float64, column []float64, colnum int) {
	for i := range table {
		column[i] = table[i][colnum]
	}
}

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
