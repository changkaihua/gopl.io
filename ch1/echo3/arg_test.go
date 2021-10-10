package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestArgs1(t *testing.T) {
	// 0.002s
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func TestArgs2(t *testing.T) {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
