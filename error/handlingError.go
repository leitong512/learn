package error

import (
	"fmt"
	"strings"
)

func xxx() error {
	return fmt.Errorf("EOF:123")
}

func handling() {
	err := xxx()
	if err != nil {
		errStr := err.Error()
		if strings.Contains(errStr,"EOF"){

		}
	}
}
//二
type MyError struct {
	Msg string
	File string
	Line int
}
func (e *MyError)Error() string {
	return fmt.Sprintf("%s:%d:%s", e.File,e.Line,e.Msg)
}
func test() error {
	return &MyError{
		Msg: "something happened",
		File: "server.go",
		Line: 42,
	}
}
func ErrorTypes() {
	err := test()
	switch err.(type) {
	case nil:
		//call succeeded, nothing to do
		case *MyError:
			fmt.Println("error occurred on line:", err.(*MyError).Line)
	default:
		//unknown err

	}
}

//三
//net包
type temporary interface {
	Temporary() bool
}
// IsTemporary returns true if err is temporary
func IsTemporary(err error) bool {
	te, ok := err.(temporary)
	return ok && te.Temporary()
}