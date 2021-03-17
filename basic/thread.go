package main

import (
	"fmt"
	"sync"
)

var wait sync.WaitGroup

func printWord(str string) {
	fmt.Printf(str)
	wait.Done()
}
func main() {
	var n int
	str := []string{"A", "B", "C", "D"}
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			return
		}
		for i := 0; i < n; i++ {
			wait.Add(1)
			go printWord(str[0])
			wait.Wait()
			wait.Add(1)
			go printWord(str[1])
			wait.Wait()
			wait.Add(1)
			go printWord(str[2])
			wait.Wait()
			wait.Add(1)
			go printWord(str[3])
			wait.Wait()
		}
		fmt.Println()
	}

}
