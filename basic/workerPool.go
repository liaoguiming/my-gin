package basic

import (
	"fmt"
	"math/rand"
)

type Job struct {
	Id      int
	RandNum int
}

type Result struct {
	job *Job
	sum int
}

/**
创建goroutine池
num 开几个协程
*/
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			//执行运算
			//遍历job管道所有数据 进行相加
			for job := range jobChan {
				rand_num := job.RandNum
				var sum int
				for rand_num != 0 {
					tmp := rand_num % 10
					sum += tmp
					rand_num /= 10
				}
				res := &Result{
					job: job,
					sum: sum,
				}
				//将结果扔给管道
				resultChan <- res
			}
		}(jobChan, resultChan)
	}
}

func main() {
	//创建两个管道
	jobChan := make(chan *Job, 128)
	resultChan := make(chan *Result, 128)
	createPool(64, jobChan, resultChan)
	go func(resultChan chan *Result) {
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id,
				result.job.RandNum, result.sum)
		}
	}(resultChan)
	var id int
	for {
		id++
		rand_num := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: rand_num,
		}
		jobChan <- job
	}
}
