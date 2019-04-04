package main

import "fmt"

type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}
	return value
}

// Вместо конструктора:
func NewTrade(symbol string, volume int, price float64, buy bool) (*Trade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol can't be empty")
	}
	if volume <= 0 {
		return nil, fmt.Errorf("volume must be >0 (was %d)", volume)
	}
	if price <= 0.0 {
		return nil, fmt.Errorf("Price must be >0 (was %f)", price)
	}
	trade := &Trade{symbol, volume, price, buy}
	return trade, nil
}

func main() {
	t1 := Trade{"MSFT", 10, 99.98, true}
	fmt.Println(t1)

	fmt.Printf("%+v\n", t1)

	fmt.Println(t1.Symbol)

	t2 := Trade{
		Symbol: "MSFT",
		Volume: 10,
		Price:  99.98,
		Buy:    true,
	}
	fmt.Printf("%+v\n", t2)

	fmt.Println(t2.Value())

	t3 := Trade{}
	fmt.Printf("%v\n", t3)

	t4, err := NewTrade("GOOG", 1, 999.0, false)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(t4)
	}
}
