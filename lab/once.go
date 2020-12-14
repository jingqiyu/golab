package lab

import (
	"fmt"
	"sync"
)

var once sync.Once
var initS *OnceInitS

type OnceInitS struct {
	A string `json:"a"`
}

func init() {
	fmt.Println("in init...")
	once.Do(func() {
		fmt.Println(" in once do")
		initS = new(OnceInitS)
		initS.A = "a"
	})
}

func OutPutInits() {
	fmt.Println(initS)
}
