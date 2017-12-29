package main

import (
	"fmt"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

func add(t *tree, val int) *tree {
	if t == nil {
		return &tree{value: val}
	}

	if val < t.value {
		t.left = add(t.left, val)
	} else {
		t.right = add(t.right, val)
	}

	return t
}

func AddValuesToTree(vals []int) *tree {
	var root *tree
	for _, val := range vals {
		root = add(root, val)
	}

	return root
}

// the string method, Stringer interface
func (t *tree) String() string {
	if t == nil {
		return ""
	}
	return t.left.String() + strconv.Itoa(t.value) + t.right.String()
	// Shit, string function does convert as expected
	//return string(t.value)
}

func main() {
	t := AddValuesToTree([]int{4, 3, 1, 7, 9})
	s := t.String()
	fmt.Println(s)
}
