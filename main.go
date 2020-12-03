package main

import (
	"fmt"
	"time"
)

var count int
var noteLeft int
var noteRight int
func main() {
	go func() {
		noteRight++
		if noteLeft == 0 {
			if count == 0 {
				fmt.Println("testright")
				count++
			}
		}
		noteRight--
	}()
	go func() {
		noteLeft++
		for noteRight== 1 {

		}
		fmt.Println("test")
		if count == 0 {
			fmt.Println("test1")
			count++
		}
		noteLeft--
	}()

	time.Sleep(time.Second)
	fmt.Println("count:",count)
}
