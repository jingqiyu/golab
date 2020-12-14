package dp

import (
	"fmt"
	"testing"
)

func TestNumArray_SumRange(t *testing.T) {
	constructor := Constructor([]int{1, 2, 3, 4, 5})
	fmt.Println(constructor.SumRange(1, 2))
}
