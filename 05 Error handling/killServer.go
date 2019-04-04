package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// func setupLogging() {
// 	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		return
// 	}
// 	log.SetOutput(out)
// }

func killServer(pidFile string) error {
	content, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return errors.Wrap(err, "Can't open file")
	}
	pid, err := strconv.Atoi(strings.TrimSpace(string(content)))
	if err != nil {
		return errors.Wrap(err, "bad pid")
	}
	if err := os.Remove(pidFile); err != nil {
		log.Printf("warn: can't remove pidfile %s", pidFile)
	}
	fmt.Printf("pid: %v (%T)", pid, pid)
	// pid := strconv.Atoi(content)
	// fmt.Printf("pid: %v (%T)", pid, pid)
	return nil
}

func main() {
	// setupLogging()
	err := killServer("a.pid")
	if err != nil {
		log.Println(err)
		// fmt.Println(err)
	}
}
