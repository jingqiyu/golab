package list

import (
	"testing"
)

func Test_getDecimalValue(t *testing.T) {
	list := GenList(1, 2, 3, 4, 5)
	println(getDecimalValue(list))
}
