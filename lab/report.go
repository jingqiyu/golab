package lab

import (
	"fmt"
	"strings"
)

var Fields = []string{
	"HuaweiFoldPushDetect", "IsIosPublish", "search_tabs",
}

func AutoGen() {
	var funcDef []string
	var call []string
	for _, v := range Fields {
		funcDef = append(funcDef, fmt.Sprintf("func (c *RemoteConfig) fill%s (r *dao.RemoteConfig) {}", strings.Title(v)))
		call = append(call, fmt.Sprintf("c.fill%s(&r)", strings.Title(v)))
	}

	for _, v := range funcDef {
		fmt.Println(v)
	}

	fmt.Println(" ---------------------- ")

	for _, v := range call {
		fmt.Println(v)
	}
}
