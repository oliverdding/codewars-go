package kata

func Snail(snaipMap [][]int) []int {
	if len(snaipMap) == 0 || len(snaipMap[0]) == 0 {
		return []int{}
	}
	result := make([]int, 0, len(snaipMap)*len(snaipMap))
	rowStart, columnStart := 0, 0
	rowSize, columnSize := len(snaipMap), len(snaipMap)
	for len(result) < cap(result) {
		row := rowStart
		column := columnStart
		for ; column-columnStart < columnSize; column += 1 {
			result = append(result, snaipMap[row][column])
		}
		row = rowStart + 1
		column = columnStart + columnSize - 1
		for ; row-rowStart < rowSize; row += 1 {
			result = append(result, snaipMap[row][column])
		}
		row = rowStart + rowSize - 1
		column = columnStart + columnSize - 2
		for ; column >= columnStart; column -= 1 {
			result = append(result, snaipMap[row][column])
		}
		row = rowStart + rowSize - 2
		column = columnStart
		for ; row > rowStart; row -= 1 {
			result = append(result, snaipMap[row][column])
		}
		rowStart, columnStart = rowStart+1, columnStart+1
		rowSize, columnSize = rowSize-2, columnSize-2
	}
	return result
}
