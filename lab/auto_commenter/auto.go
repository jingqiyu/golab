package auto_commenter

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var ignoreDirPath = map[string]int{
	"/Users/jingqiyu/dev/trpc/comment/jce": 1,
	".git":                                 1,
}

//判断目录是否存在
func isDir(dir string) bool {
	info, err := os.Stat(dir)

	if err == nil {

		return false
	}
	return info.IsDir()
}

func isGo(name string) bool {
	if !strings.Contains(name, ".") {
		return false
	}
	split := strings.Split(name, ".")
	if len(split) < 2 {
		return false
	}

	return split[len(split)-1] == "go"

}

func AutoCommenter(dirPath string, show bool) {
	files, _ := ioutil.ReadDir(dirPath)
	for _, file := range files {
		if file.IsDir() {
			if _, ok := ignoreDirPath[dirPath+"/"+file.Name()]; ok {
				continue
			}
			AutoCommenter(dirPath+"/"+file.Name(), show)
		}
		if !isGo(file.Name()) {
			continue
		}
		Comments(dirPath+"/"+file.Name(), show)
	}
	return
}

func Comments(filePath string, show bool) {
	if show {
		fmt.Println(filePath)
	}
	fi, err := os.Open(filePath)
	if err != nil {
		fmt.Println("ErrorPBFilePath PleaseCheck")
		return
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	var lines []string
	for {
		contentBytes, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		content := string(contentBytes)
		lines = append(lines, content)
	}

	var ret []string
	var lineNo int

	for _, line := range lines {
		isType, structName := isTypeString(line, lines, lineNo)
		isFunc, funcName := isFuncType(line, lines, lineNo)

		if !isType && !isFunc {
			ret = append(ret, line)
		}

		if isType {
			ret = append(ret, fmt.Sprintf("// %s \n%s", structName, line))
		} else if isFunc {
			ret = append(ret, fmt.Sprintf("// %s \n%s", funcName, line))

		}
		lineNo++
	}

	var writeBuffer bytes.Buffer

	for _, v := range ret {
		writeBuffer.WriteString(v + "\n")
	}
	_ = ioutil.WriteFile(filePath, writeBuffer.Bytes(), 600)
}

func isTypeString(line string, lines []string, lineNo int) (bool, string) {

	line = strings.TrimSpace(line)

	// 必须以type 开头的
	if strings.Index(line, "type") != 0 {
		return false, ""
	}

	// 防命名
	split := strings.Split(line, " ")

	if len(split) < 3 {
		return false, ""
	}

	if split[0] != "type" {
		return false, ""
	}

	name := split[1]
	if split[2] == "struct" {
		name += " object"
	}

	if lineNo == 1 {
		return true, name
	}

	lastLine := lines[lineNo-1]

	//已经有注释的不在重复添加注释
	if strings.Index(lastLine, "//") == 0 {
		return false, ""
	}

	return true, name
}

// func 的两种形式
// func (h ***) getXX() { 带接收器 接收器命名 ）

// 还有 func GetXXX() {

func isFuncType(line string, lines []string, lineNo int) (bool, string) {

	line = strings.TrimSpace(line)

	// 必须以func 开头的
	if strings.Index(line, "func") != 0 {
		return false, ""
	}

	split := strings.Split(line, " ")

	if len(split) < 2 {
		return false, ""
	}

	if split[0] != "func" {
		return false, ""
	}

	var receiver string
	var funcName string
	// 可能是带接收器的
	if strings.Index(split[1], "(") == 0 {
		if len(split) >= 3 && strings.Index(split[2], ")") == len(split[2])-1 {
			funcName = split[3]
			funcName = split[3][0:strings.Index(split[3], "(")]
			receiver = split[2][:len(split[2])-1]
		}
	} else {
		funcName = split[1]
		funcName = split[1][0:strings.Index(split[1], "(")]
	}

	if lineNo == 1 {
		if receiver != "" {
			return true, fmt.Sprintf("%s -> %s", receiver, funcName)
		}
	}

	lastLine := lines[lineNo-1]

	//已经有注释的不在重复添加注释
	if strings.Index(strings.TrimSpace(lastLine), "//") == 0 ||
		strings.Index(strings.TrimSpace(lastLine), "*/") == 0 {
		return false, ""
	}
	if receiver != "" {
		return true, fmt.Sprintf("%s -> %s", receiver, funcName)
	} else {
		return true, funcName
	}
}
