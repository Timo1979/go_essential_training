package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("Content-Type")
	fmt.Printf("%s -> %s\n", url, ctype)
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
	t1 := time.Now().UnixNano()
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			returnType(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
	t2 := time.Now().UnixNano()
	return t2 - t1
}
