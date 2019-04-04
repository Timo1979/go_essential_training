package main

import (
	"fmt"
)

func main() {
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Printf("loons = %v (%T)\n", loons, loons)

	fmt.Println(len(loons))

	fmt.Println("----")

	fmt.Println(loons[1])

	fmt.Println("----")
	fmt.Println(loons[1:])

	fmt.Println("---- for i:=0 ;i<len(loons);i++")
	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}

	fmt.Println("---- for i := range loons {")
	for i := range loons {
		fmt.Println(i)
	}

	fmt.Println("---- for i, name := range loons {")
	for i, name := range loons {
		fmt.Printf("%s at %d\n", name, i)
	}

	fmt.Println("---- for _, name := range loons {")
	for _, name := range loons {
		fmt.Println(name)
	}

	fmt.Println("---- append")
	loons = append(loons, "elmer")
	fmt.Println(loons)

}
