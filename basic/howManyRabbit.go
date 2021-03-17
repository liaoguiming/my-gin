package basic

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)




func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		input := scanner.Text()
		if len(input) == 0{break}
		numOfMonth, _ := strconv.Atoi(input)
		numOfOne := 0
		numOfTwo := 0
		numOfThree := 0
		for i := 0; i < numOfMonth; i++{
			if i == 0{
				numOfOne += 1
				continue
			}
			numOfThree += numOfTwo
			numOfTwo = numOfOne
			numOfOne = numOfThree
		}

		fmt.Println(numOfOne + numOfTwo + numOfThree)
	}

}
