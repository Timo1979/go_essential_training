package main

import (
	"fmt"
	"net/http"
	"time"
)

func returnType(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("Content-Type")
	return fmt.Sprintf("%s -> %s\n", url, ctype)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	var t int64 = 0
	for i := 0; i < 5; i++ {
		t += onePass(urls)
	}
	fmt.Printf("time of execution: %d\n", t/5000000)
}

func onePass(urls []string) int64 {
	ch := make(chan string)
	t1 := time.Now().UnixNano()
	for _, url := range urls {
		go func(url string) {
			ch <- returnType(url)
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-ch)
	}
	t2 := time.Now().UnixNano()
	return t2 - t1
}
