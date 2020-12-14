package dp

func minimumTotal(triangle [][]int) int {
	r := len(triangle)

	for i := 1; i < r; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][j]
				continue
			}
			if j == i {
				triangle[i][j] += triangle[i-1][j-1]
				continue
			}
			if triangle[i-1][j-1] > triangle[i-1][j] {
				triangle[i][j] += triangle[i-1][j]
			} else {
				triangle[i][j] += triangle[i-1][j-1]
			}

		}
	}
	lastR := triangle[len(triangle)-1]
	min := lastR[0]

	for i := 0; i < len(lastR); i++ {
		if lastR[i] < min {
			min = lastR[i]
		}
	}

	return min
}
