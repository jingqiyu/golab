package huisu

import (
	"fmt"
	"testing"
)

func Test_combinationSum2(t *testing.T) {
	sum2 := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	fmt.Println(sum2)
}

func Test_combinationSum3(t *testing.T) {
	sum2 := combinationSum3(3, 9)
	fmt.Println(sum2)
}
