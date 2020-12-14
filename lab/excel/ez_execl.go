package excel

import (
	"fmt"
	"reflect"
)

const (
	exTag          = "ex"
	exColsIndexTag = "ex_col"
)

type ExcelTypeStruct struct {
	Name    string `ex:"name" ex_col:"A"`
	Age     int    `ex:"age" ex_col:"B"`
	Address string `ex:"address"  ex_col:"C"`
}

func DecodeStruct(e ExcelTypeStruct) {

	var titleM = make(map[string]string)

	//rv := reflect.ValueOf(&e).Elem()
	rt := reflect.TypeOf(&e).Elem()

	field := rt.NumField()

	for i := 0; i < field; i++ {

		exTitle := rt.Field(i).Tag.Get("ex")
		exCol := rt.Field(i).Tag.Get("ex_col")

		titleM[exCol+"1"] = exTitle
	}

	fmt.Println(titleM)
}
