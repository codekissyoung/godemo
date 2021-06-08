package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
$ go run main.go data.txt
2	hello world
1	我喜欢你
1	nice to meet you
*/

var counts = make(map[string]int)

func main() {
	files := os.Args[1:]

	// 从标准输入中
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, fileName := range files {
			file, err := os.Open(fileName)
			if err != nil {
				_, err := fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				if err != nil {
					fmt.Println("Fprintf error")
				}
				continue
			}
			countLines(file, counts)
			file.Close()
		}
	}
	for line, n := range counts {
		fmt.Printf("%d	%s	\n", n, line)
	}
}

func countLines(reader io.Reader, counts map[string]int) {
	input := bufio.NewScanner(reader)
	for input.Scan() {
		if input.Text() != "" { // 不统计空行
			counts[input.Text()]++
		}
	}
}
