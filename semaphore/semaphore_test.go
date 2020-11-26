package semaphore

import (
	"fmt"
	"testing"
	"time"
)

var count = 0
func TestSemaphore(t *testing.T) {
	sem := NewSemaphore(100) //不能当锁用sem(0)
	for i :=0; i <1000; i++ {
		go func() {
			sem.Acquire()
			count++
			sem.Release()
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(count)
}
