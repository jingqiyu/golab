package lab

import (
	"fmt"
	"os"
)

func GetEnv() string {

	getenv := os.Getenv("ENT_TYPE")
	fmt.Println(getenv)

	return getenv
}
