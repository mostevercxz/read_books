package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int

func (wc *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	*wc += WordCounter(count)
	return count, nil
}

type LineCounter int

func (lc *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	*lc += LineCounter(count)
	return count, nil
}

func main() {
	var wc WordCounter
	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
	var lc LineCounter
	wc.Write([]byte(input))
	lc.Write([]byte(input))
	fmt.Println(wc)
	fmt.Println(lc)
}
