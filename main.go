package main

import (
	"fmt"

	"github.com/khatibomar/play-currency/forex"
)

type USD struct{}

func (u USD) Code() string {
	return "USD"
}

func (u USD) Symbol() string {
	return "$"
}

func (u USD) Ratio() float64 {
	return 100
}

func main() {
	a := forex.NewAmount[USD](100)
	fmt.Println(a)
	b := forex.NewAmount[USD](200)
	a.Add(b)
	fmt.Println(a)
}
