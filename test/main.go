package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var count int32

func main() {
	i := int32(0)
	for ; i < 10; i++ {
		wg.Add(1)
		go func(i int32) {
			defer func() {
				wg.Done()
			}()
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}
	wg.Wait()
}
func trigger(i int32, fn func()) {
	for {
		if n := atomic.LoadInt32(&count); n == i {
			fn()
			atomic.AddInt32(&count, 1)
			break
		}
		time.Sleep(time.Second)
	}
}
