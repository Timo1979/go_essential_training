package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("Content-Type")
	if ctype == "" {
		return "", fmt.Errorf("can't find content-type header")
	}
	return ctype, nil
}

func main() {
	// t, err := contentType("http://www.gto.by")
	t, err := contentType("http://localhost:8081/pays/s/open/ownDsList")
	if err != nil {
		fmt.Println("Ошибка")
	} else {
		fmt.Println(t)
	}
}
