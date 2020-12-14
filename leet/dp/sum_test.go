package dp

import (
	"fmt"
	"testing"
)

func Test_prefixSumArr(t *testing.T) {
	prefixSumArr([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}})

}

func Test_matrixBlockSum(t *testing.T) {
	fmt.Println(matrixBlockSum([][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
		[]int{7, 8, 9},
		[]int{7, 8, 9},
	}, 3))

}
