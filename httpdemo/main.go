package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var countMtx sync.Mutex
var count int

func countReq() {
	countMtx.Lock()
	defer countMtx.Unlock()
	count++
}

func index(w http.ResponseWriter, r *http.Request) {
	countReq()
	_, _ = fmt.Fprintf(w, "Hello World!")
}
func counter(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Count : %d \n", count)
}
func timeNow(w http.ResponseWriter, r *http.Request) {
	countReq()
	t := time.Now()
	timeStr := fmt.Sprintf("{\"time\":\"%s\"}", t)
	_, _ = w.Write([]byte(timeStr))
}

// curl "localhost:8080/echo?name=link&age=23"
// GET /echo?name=link&age=23 HTTP/1.1
// Host : "localhost:8080"
// RemoteAddr : "127.0.0.1:43606"
// Header["User-Agent"] = ["curl/7.58.0"]
// Header["Accept"] = ["*/*"]
// Form["name"] = ["link"]
// Form["age"] = ["23"]
func echo(w http.ResponseWriter, r *http.Request) {
	countReq()
	_, _ = fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
	_, _ = fmt.Fprintf(w, "Host : %q \n", r.Host)
	_, _ = fmt.Fprintf(w, "RemoteAddr : %q \n", r.RemoteAddr)
	for k, v := range r.Header {
		_, _ = fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	for k, v := range r.Form {
		_, _ = fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/time", timeNow)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/echo", echo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
