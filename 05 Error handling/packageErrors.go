package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

type Config struct {
	//
}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "Can't open config file")
	}
	defer file.Close()
	cfg := &Config{}
	return cfg, nil
}

func setupLogging() {
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	log.SetOutput(out)
}

func main() {
	setupLogging()
	_, err := readConfig("ccccc")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		log.Printf("Error: %+v\n", err)
		os.Exit(1)
	}
}
