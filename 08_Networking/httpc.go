// Makign HTTP calls
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Job struct {
	User   string `json:"user"`
	Action string `json:"action"`
	Count  int    `json:"count"`
}

func main() {
	// GET request
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalf("error: can't call httpbin.org")
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)

	fmt.Println("######")

	job := &Job{"Saitama", "punch", 1}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(job); err != nil {
		log.Fatalf("Error cant encode job - %s", err)
	}
	fmt.Printf("buf to post: %+v\n", buf.String())
	resp, err = http.Post("https://httpbin.org/post", "application/json", &buf)
	if err != nil {
		log.Fatalf("can't send request to httpbin - %s", err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
