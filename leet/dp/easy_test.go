package dp

import (
	"fmt"
	"testing"
)

func Test_maxValue(t *testing.T) {
	value := maxValue([][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	})
	fmt.Println(value)
}
