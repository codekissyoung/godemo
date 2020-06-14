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

	n, err := io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Read", n, "Bytes")

	if err := resp.Body.Close(); err != nil {
		log.Fatalln(err)
	}

}
