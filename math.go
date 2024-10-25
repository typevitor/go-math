package main

import (
	"fmt"
	"errors"
)

func main() {
	fmt.Println(Sum(10,10))
}

func Sum(a int, b int) int {
	return a + b
}

func Sub(a int, b int) int {
	return a - b
}

func Times(a int, b int) int {
	return a * b
}

func Div(a int, b int) (float32, error) {
	if b == 0 {
		return 0, errors.New("Divisor não pode ser zero")
	}

	return float32(a) / float32(b), nil
}