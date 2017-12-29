package main

import (
	"fmt"
	"io"
	"os"
)

type countingWriter struct {
	io.Writer
	count int64
}

func (cw *countingWriter) Write(p []byte) (int, error) {
	count, err := cw.Writer.Write(p)
	cw.count += int64(count)
	return count, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var cw countingWriter
	cw.Writer = w
	return &cw, &cw.count
}

func main() {
	cw, countptr := CountingWriter(os.Stdout)
	cw.Write([]byte("hello"))
	cw.Write([]byte("\ngo\n"))
	fmt.Printf("write %d bytes", *countptr)
}
