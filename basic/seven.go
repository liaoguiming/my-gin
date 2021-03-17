package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	for {
		sum := 0
		_, err := fmt.Scan(&n)
		if err != nil{
			return
		}
		for i := 7; i <= n; i++ {
			//能被7整除
			if i%7 == 0 || strings.Index(strconv.Itoa(i), "7") != -1 {
				sum++
				continue
			}
		}
		fmt.Println(sum)
	}
}
