package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type MyReader struct {
	io.Reader
	limit int64
}

func (r *MyReader) Read(p []byte) (int, error) {
	count, err := r.Reader.Read(p)
	if count > int(r.limit) {
		p = p[:r.limit]
		return int(r.limit), io.EOF
	} else {
		return count, err
	}
}

func LimitReader(r io.Reader, n int64) io.Reader {
	var reader MyReader
	reader.limit = n
	reader.Reader = r
	return &reader
}

func main() {
	reader := strings.NewReader("hello,go~")
	limitReader := LimitReader(reader, 5)
	buf := new(bytes.Buffer)
	buf.ReadFrom(limitReader)
	fmt.Println(buf)
}
