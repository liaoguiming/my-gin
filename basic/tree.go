package basic

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	num := 10
	ch1, ch2 := make(chan int, num), make(chan int, num)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < num; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	//测试内容
	t1 := tree.New(10)
	t2 := tree.New(10)
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(Same(t1, t2))
}
