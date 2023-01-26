package main

import (
	"mymodule/Math"
	"mymodule/Print"
)

func main () {
	a := 3
	b := 4
	sl := []int{1, 2, 3, 4}
	Math.Sum(a, b)
	Print.Print(sl)
}