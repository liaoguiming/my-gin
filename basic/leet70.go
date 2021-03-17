package basic

import "fmt"

//该题考察了斐波那契e数列的应用 第N个台阶 肯定是由N-1 和 N-2想加得到的
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b, temp := 1, 2, 0
	for i := 3; i <= n; i++ {
		temp = a + b
		a = b
		b = temp
	}
	return temp
}
func main() {
	a := 3
	fmt.Println(climbStairs(a))
}
