package currency

type Currency interface {
	Code() string
	Symbol() string
	Ratio() float64
}

type Amount[T Currency] struct {
	Value int64
}

func NewAmount[T Currency](value int64) Amount[T] {
	return Amount[T]{
		Value: value,
	}
}

func (a *Amount[T]) Add(amount Amount[T]) {
	a.Value += amount.Value
}
