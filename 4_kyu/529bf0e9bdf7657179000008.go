package kata

func ValidateSolution(m [][]int) bool {
	for r := 0; r < 9; r++ {
		row, col, box := 0, 0, 0
		for c := 0; c < 9; c++ {
			row ^= 1 << uint(m[r][c])
			col ^= 1 << uint(m[c][r])

			i := (r%3)*3 + c%3
			j := (r/3)*3 + c/3
			box ^= 1 << uint(m[j][i])
		}
		if row != 1022 || col != 1022 || box != 1022 {
			return false
		}
	}
	return true
}
