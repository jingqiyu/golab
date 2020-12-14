package excel

import (
	"testing"
)

func TestDecodeStruct(t *testing.T) {
	DecodeStruct(ExcelTypeStruct{
		Name:    "",
		Age:     0,
		Address: "",
	})
}
