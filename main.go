package main

import "fmt"

type IceCreamMaker interface {
	Hello()
}

type Ben struct {
	name string
}

func (b *Ben) Hello() {
	if b.name != "Ben" {
		fmt.Printf("Ben says, \"Hello my name is %s\"\n", b.name)
	}
}

type Jerry struct {
	field2 int
	field1 *[5]byte
}

func (j *Jerry) Hello() {
	name := string((*j.field1)[:])
	if name != "Jerry" {
		fmt.Printf("Jerry says, \"Hello my name is %s\"\n", name)
	}
}
func main() {
	var ben = &Ben{"Ben"}
	var jerry = &Jerry{5, &[5]byte{'J', 'e', 'r', 'r', 'y'}}
	var maker IceCreamMaker = ben

	var loop0, loop1 func()
	loop0 = func() {
		maker = ben
		go loop1()
	}
	loop1 = func() {
		maker = jerry
		go loop0()
	}
	fmt.Printf("Ben: %p Jerry: %p\n", ben, jerry)
	go loop0()
	for {
		maker.Hello()
	}
}