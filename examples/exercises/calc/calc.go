package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrDivisionByZero  = errors.New("division by zero")
	ErrUnknownOperator = errors.New("unknown operator")
)

func calc(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, ErrDivisionByZero
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("%w: %s", ErrUnknownOperator, op)
	}
}

func main() {
	cases := []struct {
		a, b float64
		op   string
	}{
		{10, 2, "/"},
		{10, 0, "/"},
		{3, 4, "+"},
		{3, 4, "?"},
	}

	for _, c := range cases {
		result, err := calc(c.a, c.b, c.op)
		if err != nil {
			log.Printf("calc(%v, %v, %q) - ошибка: %v", c.a, c.b, c.op, err)
		} else {
			fmt.Printf("calc(%v, %v, %q) - %v\n", c.a, c.b, c.op, result)
		}
	}
}
