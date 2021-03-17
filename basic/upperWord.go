package basic

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		sum := 0
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		for i := 0; i < len(str); i++ {
			if str[i] >= 65 && str[i] <= 90 {
				sum++
			}
		}
		fmt.Println(sum)
	}
}
