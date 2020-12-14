package stack

import (
	"fmt"
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

func Test_finalPrices(t *testing.T) {
	fmt.Println(finalPrices([]int{10, 1, 1, 6}))
}
