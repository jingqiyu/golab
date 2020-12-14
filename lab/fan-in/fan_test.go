package fan_in

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	in := producer(100)

	// FAN-OUT
	maxRoutines := 6

	squares := make([]<-chan squareResult, maxRoutines)
	for i := 0; i < maxRoutines; i++ {
		squares[i] = square(in)
	}

	/*c1 := square(in)
	c2 := square(in)
	c3 := square(in)
	c4 := square(in)
	c5 := square(in)
	c6 := square(in)
	// consumer
	for v := range merge(c1, c2, c3, c4, c5, c6) {
		fmt.Println(v)
	}*/

	var result []squareResult
	for v := range merge(squares...) {
		result = append(result, v)
	}
	fmt.Println(len(result))
	fmt.Println(result)
}
