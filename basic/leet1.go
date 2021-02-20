package main

import "fmt"

func twoSum(nums []int, target int) []int {
	arr := []int{0, 0}
	res := false
	for i := 0; i < len(nums)-1; i++ {
		nowNum := target - nums[i]
		for j := i+1; j < len(nums); j++ {
			if nowNum-nums[j] == 0 {
				arr[0] = i
				arr[1] = j
				res = true
				break
			}
		}
		if res {
			break
		}
	}
	return arr
}

func main() {
	a := []int{2, 5, 5, 11}
	fmt.Println(twoSum(a, 10))
}
