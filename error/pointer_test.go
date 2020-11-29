package error

import (
	"errors"
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	var ErrNameType = New("EOF")
	var ErrStructType = errors.New("EOF")

	if ErrNameType == New("EOF") {
		fmt.Println("Named Type error")
	}
	if ErrStructType == errors.New("EOF") {
		fmt.Println("Struct Type Error")
	}
}