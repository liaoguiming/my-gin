package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseStr(str string) string {
	newStr := ""
	strSplit := strings.Split(str, "")
	for i := len(strSplit) - 1; i >= 0; i-- {
		newStr = newStr + strSplit[i]
	}
	return newStr
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := input.Text()
	newStr := reverseStr(str)
	fmt.Println(newStr)
}
