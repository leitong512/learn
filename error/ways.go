package error

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

var path = "/c/a"

// Indented flow is for errors
func demo1() {
	_, err := os.Open(path)
	if err != nil {
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

func authenticate(a string) error {
	return errors.New("1")
}

//统计io.Reader读取内容的行数
func CountLinesOne(r io.Reader) (int, error) {
	var (
		br    = bufio.NewReader(r)
		lines int
		err   error
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
	return lines, nil
}
func CountLinesTwo(r io.Reader) (int, error) {
	var (
		sc    = bufio.NewScanner(r)
		lines int
	)
	for sc.Scan() {
		lines++
	}
	return lines, sc.Err()
}

//三
type Header struct {
	Key, Value string
}
type Status struct {
	Code   int
	Reason string
}

func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
	_, err := fmt.Fprintf(w, "Http/1.1 %d %s\r", st.Code, st.Reason)
	if err != nil {
		return err
	}
	for _, h := range headers {
		_, err := fmt.Fprintf(w, "%s:%s\r\n", h.Key, h.Value)
		if err != nil {
			return err
		}
	}
	if _, err := fmt.Fprint(w, "\r\n"); err != nil {
		return err
	}
	_, err = io.Copy(w, body)
	return err
}

//减少代码的
type errWriter struct {
	io.Writer
	err error
}

func (e *errWriter) Write(buf []byte) (int, error) {
	if e.err != nil {
		return 0, e.err
	}
	var n int
	n, e.err = e.Writer.Write(buf)
	return n, nil
}

func WriteResponseTwo(w io.Writer, st Status, headers []Header,
	body io.Reader) error {
	ew := &errWriter{Writer: w}
	fmt.Fprintf(ew, "Http/1.1 %d %s\r\n", st.Code, st.Reason)
	for _, h := range headers {
		fmt.Fprintf(w, "%s:%s\r\n", h.Key, h.Value)
	}
	fmt.Fprint(ew, "\r\n")
	io.Copy(ew, body)
	return ew.err

}
