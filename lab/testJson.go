package lab

import (
	"encoding/json"
	"fmt"
)

func TestJson() {
	type A struct {
		A string `json:"a"`
		B string `json:"-"`
	}

	A1 := A{A: "Abc", B: "ac"}
	marshal, _ := json.Marshal(A1)
	fmt.Println(string(marshal))
}
