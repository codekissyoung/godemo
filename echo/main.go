package main

import (
	"fmt"
	"os"
	"strings"
)

// slice s[m:n] 其中 0 <= m <= n <= len(s) , 包含 n - m 个元素
func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println("Join : ", strings.Join(os.Args[1:], ","))
}
