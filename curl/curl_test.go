package main

import (
	"net/http"
	"strconv"
	"testing"
)

const checkMarkRight = "\u2713"
const checkMarkFault = "\u2717"

// go test -v
func TestDownload(t *testing.T) {
	urls := []struct {
		url        string
		statusCode int
	}{
		{"http://blog.codekissyoung.com",
			200,
		},
		{"http://baidu.com",
			200,
		},
	}
	for _, u := range urls {
		resp, err := http.Get(u.url)
		if err != nil {
			t.Errorf("Should be able to Get the url %v", checkMarkFault)
			return
		}
		if resp.StatusCode == u.statusCode {
			t.Logf("Can get the url %v", checkMarkRight)
		} else {
			t.Errorf("Can not get the url %v", checkMarkFault)
		}
		resp.Body.Close()
	}
}

// go test -run="none" -bench=. -benchtime="1s" -benchmem -v
func BenchmarkFormat(b *testing.B) {
	number := int64(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(number, 10)
	}
}

func BenchmarkItoa(b *testing.B) {
	number := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}
}
