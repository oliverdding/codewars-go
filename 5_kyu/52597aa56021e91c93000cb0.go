package kata

func GetNonZeroIndex(arr []int) int {
	for i, v := range arr {
		if v != 0 {
			return i
		}
	}
	return len(arr) // Special
}

/*
Solution1: BubbleSwip
Solution2: SelectionSwip

	0 0 -> nonZeroIdx+GetNonZeroIndex(arr[nonZeroIdx:])
	0 1 -> arr[curIdx] = arr[nonZeroIdx], arr[nonZeroIdx] = 0, curIdx+1 nonZeroIdx+1
	1 0 -> curIdx+1 nonZeroIdx+GetNonZeroIndex(arr[nonZeroIdx:])
	1 1 -> curIdx+1 nonZeroIdx+1

Below is selection swip.
*/
func MoveZeros(arr []int) []int {
	curIdx := 0
	nonZeroIdx := 0
	for nonZeroIdx < len(arr) {
		if arr[curIdx] == 0 {
			if arr[nonZeroIdx] == 0 {
				nonZeroIdx += GetNonZeroIndex(arr[nonZeroIdx:])
			} else {
				arr[curIdx] = arr[nonZeroIdx]
				arr[nonZeroIdx] = 0
				curIdx++
				nonZeroIdx++
			}
		} else {
			curIdx++
			if arr[nonZeroIdx] == 0 {
				nonZeroIdx += GetNonZeroIndex(arr[nonZeroIdx:])
			} else {
				nonZeroIdx++
			}
		}
	}
	return arr
}
