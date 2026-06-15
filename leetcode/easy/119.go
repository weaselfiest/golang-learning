package easy

func getRow(rowIndex int) []int {
	row := []int{1}

	for i := 1; i <= rowIndex; i++ {
		newRow := make([]int, i+1)

		newRow[0] = 1
		newRow[i] = 1

		for j := 1; j < i; j++ {
			newRow[j] = row[j-1] + row[j]
		}

		row = newRow
	}

	return row
}
