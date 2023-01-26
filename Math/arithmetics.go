package Math

import "fmt"

func Sum(a int, b int) int {
	fmt.Println(a + b)
	return a + b
}

func Difference(a int, b int) int {
	fmt.Println(a - b)
	return a - b
}

func Multiplication(a int, b int) int {
	fmt.Println(a * b)
	return a * b
}

func Division(a int, b int) float64 {
	fmt.Println(float64(a)/(float64(b)))
	return float64(a)/(float64(b))
}