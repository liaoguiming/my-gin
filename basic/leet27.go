package basic

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	index := 0
	for ; index < len(nums); {
		//如果存在
		if nums[index] == val {
			//新生成不包含该值的切片
			nums = append(nums[:index], nums[index+1:]...)
			continue
		}
		index++
	}
	return len(nums)
}

func main() {
	arr := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(arr, val))
}
