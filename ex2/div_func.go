package ex1

import (
	"errors"
)

var ErrDivByZero = errors.New("division by zero")

func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivByZero

	}
	return a / b, nil
}
