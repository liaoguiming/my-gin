package basic

import "fmt"

//最小公约数可以递归寻找
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	if a < b {
		a, b = b, a
	}
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println((a * b) / gcd(a, b))
}
