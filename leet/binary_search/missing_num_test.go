package binary_search

import (
	"fmt"
	"testing"
)

func Test_missingNumber(t *testing.T) {
	src := []int{0, 1, 3, 4, 5, 6, 7, 8, 9}
	number := missingNumber(src)
	fmt.Println(number)

	src = []int{0, 1}
	number = missingNumber(src)
	fmt.Println(number)
}
