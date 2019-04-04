package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"strings"
)

type CheckResult struct {
	file string
	err  error
}

func main() {
	data, err := ioutil.ReadFile("md5sum.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	// fmt.Printf("%v\n", strings.Split(string(data), "\n"))
	counter := 0
	ch := make(chan *CheckResult)
	for _, line := range strings.Split(string(data), "\n") {
		x := strings.Fields(line)

		if len(x) == 2 {
			go func(md5sum string, file string) {
				data, err = ioutil.ReadFile(file)
				if err != nil {
					ch <- &CheckResult{file, err}
					return
				}
				calcMD5sum := fmt.Sprintf("%x", md5.Sum(data))
				if calcMD5sum != md5sum {
					ch <- &CheckResult{file, fmt.Errorf("%s: checksum error: actual %s written %s", file, calcMD5sum, md5sum)}
				} else {
					ch <- &CheckResult{file, nil}
				}
			}(x[0], x[1])
			counter++
		}
	}
	for i := 0; i < counter; i++ {
		result := <-ch
		if result.err != nil {
			fmt.Printf("%s %v\n", result.file, result.err)
		}
	}
}
