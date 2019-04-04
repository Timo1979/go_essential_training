package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type User struct {
	Name        string `json:"login"`
	PublicRepos int    `json:"public_repos"`
}

func main() {
	loginName := os.Args[1]
	resp, err := http.Get("https://api.github.com/users/" + loginName)
	if err != nil {
		log.Fatalf("Cant get info for %s - %s", loginName, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("error HTTP code %d", resp.StatusCode)
	}
	dec := json.NewDecoder(resp.Body)
	user := &User{}
	if err := dec.Decode(user); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}
	fmt.Printf("%+v\n", user)
}
