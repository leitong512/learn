package error

import "fmt"

func handle()(int, error){
	return 1,nil
}

func Dis() {
	i, err := handle()
	if err != nil {
		return
	}
	fmt.Println(i)
}