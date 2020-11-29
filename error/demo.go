package error

import (
	"errors"
	"fmt"
)

func PositiveOne(n int) bool {
	return n > -1
}

func CheckOne(n int) {
	if PositiveOne(n) {
		fmt.Println("is positive")
	} else {
		fmt.Println("is negative")
	}
}
//Positive returns true if the number is positive,false if it is negative.
//The second return value indicates if the result is valid, which in the case
// of n == 0, is not valid.
func PositiveTwo(n int) (bool,bool){
	if n == 0 {
		return false,false
	}
	return n > -1, true
}
func CheckTwo(n int) {
	pos, ok := PositiveTwo(n)
	if !ok {
		fmt.Println("is neither")
		return
	}
	if pos {
		fmt.Println("is positive")
	} else {
		fmt.Println("is negative")
	}
}
//Positive returns true if the number is positive, false if it is negative.
func PositiveThree(n int)(bool, error) {
	if n == 0 {
		return false, errors.New("undefined")
	}
	return n > -1, nil
}
func CheckThree(n int) {
	pos, err := PositiveThree(n)
	if err != nil {
		fmt.Println("is neither")
		return
	}
	if pos {
		fmt.Println("is positive")
	} else {
		fmt.Println("is negative")
	}
}