package forex

import "fmt"

type Currency interface {
	Code() string
	Symbol() string
	Ratio() float64
}

type Amount[T Currency] struct {
	Value    int64
	Currency T
}

func NewAmount[T Currency](value int64, currency T) Amount[T] {
	return Amount[T]{
		Value:    value,
		Currency: currency,
	}
}

func (a *Amount[T]) Add(amount Amount[T]) {
	a.Value += amount.Value
}

func (a Amount[T]) String() string {
	return fmt.Sprintf("%s %.2f", a.Currency.Symbol(), float64(a.Value)/a.Currency.Ratio())
}
