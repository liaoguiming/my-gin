package main

import "fmt"

func main() {
	n := 0
	//两个空瓶可以找老板借一瓶喝完凑齐三个换一瓶还给老板 所有直接除以2即可,如果只剩下一瓶 舍弃取整
	for {
		_, _ = fmt.Scan(&n)
		if n == 0 {
			break
		} else {
			times := n / 2
			fmt.Println(times)
		}
	}

}
