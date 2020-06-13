package main

import (
	"flag"
	"fmt"
	"strings"
)

/*
echoflags -h
	Usage of echoflags:
	-n	不换行 (default true)
	-s string
	指定分割府 (default " ")
*/
var n = flag.Bool("n", true, "不换行")      // 默认换行
var sep = flag.String("s", " ", "指定分割府") // 默认分割符：空格

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if *n {
		fmt.Println()
	}
}
