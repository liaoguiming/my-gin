package basic

import "fmt"

func in(target int, str_array []int) bool {
	for _, element := range str_array {
		if target == element {
			return true
		}
	}
	return false

}

func main() {
	fmt.Println(10 << 5)
	day31 := []int{1, 3, 5, 7, 8, 10, 12}
	day30 := []int{4, 6, 9, 11}
	for {
		var year, month, day int
		sum := 0
		_, err := fmt.Scan(&year, &month, &day)
		if err != nil {
			break
		}
		for i := 1; i <= month-1; i++ {
			if in(i, day31) {
				sum += 31
			} else if in(i, day30) {
				sum += 30
			} else if (year%4 == 0 && year%100 != 0) || (year%100 == 0 && year%400 == 0) {
				sum += 29
			} else {
				sum += 28
			}
		}
		sum += day
		fmt.Println(sum)
	}
}
