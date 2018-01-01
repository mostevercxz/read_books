package main

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type MyReader struct {
	content []byte
}

func (r *MyReader) Read(p []byte) (int, error) {
	copy(p, r.content)
	return len(r.content), io.EOF
}

func NewReader(content string) io.Reader {
	var reader MyReader
	reader.content = []byte(content)
	return &reader
}

func main() {
	s := `<html><body><h1>hello</h1></body></html>`
	doc, err := html.Parse(NewReader(s))
	if err != nil {
		fmt.Println("parse string error=", err)
	}
	visit(doc)
	fmt.Println("vim-go")
}

func visit(n *html.Node) {
	fmt.Println(n.FirstChild.LastChild.FirstChild.FirstChild.Data)
}
