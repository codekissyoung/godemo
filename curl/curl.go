package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	dest := io.MultiWriter(os.Stdout, file) // 同时写入标准输出 与 文件

	n, err := io.Copy(dest, resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Read", n, "Bytes")

	if err := resp.Body.Close(); err != nil {
		log.Fatalln(err)
	}

}
