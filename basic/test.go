package basic

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var x int64
var lock sync.Mutex

func hello(i int) {
	defer wg.Done()
	fmt.Println("hello goroutine!", i)
}

func add() {
	for i := 0; i < 10000; i++ {
		//lock.Lock()
		x = x + 1
		//lock.Unlock()
	}
	wg.Done()
}
func main() {
	wg.Add(3)
	go add()
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
