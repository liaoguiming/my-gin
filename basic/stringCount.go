package basic
import "fmt"
import "os"
import "bufio"

func main()  {
	assicMax()
}

func assicMax()  {
	r := bufio.NewReader(os.Stdin)
	for{

		btys, _, _ := r.ReadLine()
		if len(btys) == 0{
			break
		}
		// 先将数据读入到此数组中
		arr := make([]int, 256)

		for _, b := range btys {
			arr[b]++
		}

		// 遍历数组，每次都找一个最大的值
		for {
			max := 0
			for j := 0; j < len(arr); j++ {
				if arr[j] == 0 {
					continue
				}
				if arr[j] > arr[max] {
					max = j
				}
			}
			if arr[max] == 0 {
				break
			}
			arr[max] = 0
			fmt.Printf("%c", max)
		}
		fmt.Println()
	}
}