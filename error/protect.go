package error

import "fmt"

//我们写出去的
func Go(x func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("xxx")
			}
		}()
		x()
	}()
}

//对于别人请求，中间件底层有recover 娄底机制，不会导致我们进程挂掉，我们就返回500+
