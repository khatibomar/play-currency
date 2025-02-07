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

type GBP struct{}

func (u GBP) Code() string {
	return "GBP"
}

func (u GBP) Symbol() string {
	return "£"
}

func (u GBP) Ratio() float64 {
	return 100
}

func main() {
	a := forex.NewAmount(100, USD{})
	fmt.Println(a)
	b := forex.NewAmount(200, USD{})
	a.Add(b)
	fmt.Println(a)
	// this will fail because different currencies, uncomment and see error
	// a.Add(forex.NewAmount[GBP](50, GBP{}))
}
