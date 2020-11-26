package semaphore

import "time"

type Semaphore struct {
	sem int
	ch chan int
}

func NewSemaphore(sem int) *Semaphore {
	return &Semaphore{
		sem: sem,
		ch : make(chan int, sem),
	}
}
//获得许可
func (s *Semaphore) Acquire() {
	s.ch <- 0
}
//释放许可
func (s *Semaphore) Release() {
	<-s.ch
}
//尝试获得许可
func (s *Semaphore)TryAcquire() bool {
	select {
	case s.ch <- 0:
		return true
	default:
		return false
	}
}
//尝试指定时间内获取许可
func (s *Semaphore)TryAcquireOnTime(timeout time.Duration) bool {
	select {
	case s.ch <- 0:
		return true
	case <-time.After(timeout):
		return false
	}
}
//当前许可可用数量
func (s *Semaphore)AvailableSem()int {
	return s.sem - len(s.ch)
}