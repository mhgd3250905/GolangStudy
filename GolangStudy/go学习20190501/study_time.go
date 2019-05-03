package main

import (
	"fmt"
	"time"
)

func test() {
	time.Sleep(time.Microsecond * 100)
}

func main() {
	now := time.Now()
	fmt.Printf(now.Format("2006/01/02 15:04:05\n"))

	start := time.Now().UnixNano()
	test()
	end := time.Now().UnixNano()

	fmt.Printf("cost:%d us\n", (end-start)/1000)

}
