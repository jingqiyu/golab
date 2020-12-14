package lab

import (
	"fmt"
)

type P struct {
	P1 string `json:"p_1"`
}

type Q struct {
	P *P
}

func Point() {
	q := new(Q)
	q.P = new(P)
	fmt.Println(q.P == nil)
}
