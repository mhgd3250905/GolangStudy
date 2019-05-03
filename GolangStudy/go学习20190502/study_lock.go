package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var lock sync.Mutex
var rwLock sync.RWMutex

func main() {
	testLock()
}

func testLock() {
	var a map[int]int
	a = make(map[int]int)
	var count int32

	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10
	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			rwLock.Lock()
			b[1] = rand.Intn(100)
			rwLock.Unlock()
		}(a)
	}

	for i := 0; i < 100; i++ {
		go func(b map[int]int) {
			for {
				rwLock.RLock()
				//fmt.Println(b)
				rwLock.RUnlock()
				atomic.AddInt32(&count, 1)
			}
		}(a)
	}

	time.Sleep(10 * time.Second)

	//lock.Lock()
	//fmt.Println(a)
	fmt.Println(atomic.LoadInt32(&count))
	//lock.Unlock()

}
