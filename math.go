package main

import (
	"fmt"
	"errors"
	"math"
)

func main() {
	fmt.Println(PowInt(2,5))
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

func PowInt(a int, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func Div(a int, b int) (float32, error) {
	if b == 0 {
		return 0, errors.New("Divisor não pode ser zero")
	}

	return float32(a) / float32(b), nil
}

func Sum2(a int, b int) int {
	return a + b
}

func Sub2(a int, b int) int {
	return a - b
}

func Times2(a int, b int) int {
	return a * b
}

func PowInt2(a int, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func Div2(a int, b int) (float32, error) {
	if b == 0 {
		return 0, errors.New("Divisor não pode ser zero")
	}

	return float32(a) / float32(b), nil
}