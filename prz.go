package prz

func Columnf64(table [][]float64, column []float64, colnum int) {
	for i := range table {
		column[i] = table[i][colnum]
	}
}
