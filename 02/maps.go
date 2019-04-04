package main

import (
	"fmt"
)

func main() {

	stocks := map[string]float64{
		"AMZN": 1699.8,
		"GOOG": 1129.19,
		"MSFT": 98.61,
	}
	fmt.Println(stocks)
	stocks["MSFT"] = 12.1
	fmt.Println(stocks)

	fmt.Println(len(stocks))
	fmt.Println(stocks["MSFT"])
	fmt.Println(stocks["MSF"])

	value, ok := stocks["MSF"]
	if !ok {
		fmt.Println("MSF not found")
	} else {
		fmt.Println(value)
	}
}
