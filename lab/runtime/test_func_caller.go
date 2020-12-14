package runtime

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func a() {
	log()
}

func log() {
	pure()
}

func pure() {
	pc, filePath, line, _ := runtime.Caller(2)
	_, fileName := filepath.Split(filePath)
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(line, fileName, funcName)
}
