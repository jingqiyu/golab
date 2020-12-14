package game

import (
	"fmt"
)

func ExpUp(level int) float64 {
	var every float64 = 0.05

	if level <= 0 {
		return 0
	}

	var origin float64 = 1
	for i := 0; i < level; i++ {
		origin += origin * every
	}

	fmt.Println(origin)
	return origin - 1
}
