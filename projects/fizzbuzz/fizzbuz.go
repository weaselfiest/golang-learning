package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) string {
	var str string

	if n%3 == 0 {
		str += "Fizz"
	}
	if n%5 == 0 {
		str += "Buzz"
	}

	if str == "" {
		str = strconv.Itoa(n)
	}

	return str
}

func main() {
	for _, n := range []int{1, 3, 5, 15, 30} {
		fmt.Printf("fizzBuzz(%d) → %s\n", n, fizzBuzz(n))
	}
}
