package error

import (
	"bufio"
	"errors"
	"io"
	"os"
)
var path = "/c/a"

// Indented flow is for errors
func demo1() {
	_, err := os.Open(path)
	if err !=nil {
		//handle err
	}
	//do stuff
}
func demo2() {
	_, err := os.Open(path)
	if err == nil {
		//do stuff
	}
	//handle err
}

//Eliminate error handling by Eliminate errors
type Request struct {
	User string
}
func AuthenticateRequestOne(r Request) error {
	err := authenticate(r.User)
	if err != nil {
		return err
	}
	return nil
}

func AuthenticateRequestTwo(r Request) error {
	return authenticate(r.User)
}

func authenticate( a string) error {
	return errors.New("1")
}

//统计io.Reader读取内容的行数
func CountLinesOne(r io.Reader) (int, error) {
	var (
		br = bufio.NewReader(r)
		lines int
		err error
	)
	for {
		_, err = br.ReadString('\n')
		lines++
		if err != nil {
			break
		}
	}
	if err != io.EOF {
		return 0, err
	}
	return lines,nil
}
func CountLinesTwo(r io.Reader) (int, error) {
	var (
		sc = bufio.NewScanner(r)
		lines int
	)
	for sc.Scan() {
		lines++
	}
	return lines, sc.Err()
}