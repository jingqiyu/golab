package dp

import (
	"fmt"
	"testing"
)

func Test_minimumTotal(t *testing.T) {
	total := minimumTotal([][]int{
		[]int{2},
		[]int{3, 4},
		[]int{6, 5, 7},
		[]int{4, 1, 8, 3},
	})

	fmt.Println(total)
}
