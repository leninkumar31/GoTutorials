package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/JonCSykes/tightrope"
)

const workerCount = 10
const maxWorkerBuffer = 2000

func main() {
	work := make(chan tightrope.Request)
	fmt.Printf("type of `c` is %T\n", work)
	fmt.Printf("value of `c` is %v\n", work)
	go CreateWork(work)
	tightrope.InitBalancer(workerCount, maxWorkerBuffer, ExecuteTask).Balance(work, false, time.Duration(30)*time.Second)
}

//CreateWork....
func CreateWork(request chan tightrope.Request) {
	response := make(chan interface{})

	for {
		input := int(rand.Int31n(90))
		request <- tightrope.Request{input, response}
		output := <-response
		fmt.Println(input, output)
	}
}

//ExecuteTask....
func ExecuteTask(request tightrope.Request) {
	request.Response <- math.Sin(float64(request.Data.(int)))
	time.Sleep(time.Duration(rand.Int63n(int64(time.Second * 1))))
}
