package main

import (
	"fmt"
)

func main() {
	var n float64
	var y float64
	var x int
	fmt.Scan(&n)
	for y = n; 1 < n; n = n / 2 {
		x = x + 1
	}
	if n == 1 {
		fmt.Print(y, " is 2 of the power ", x)
	} else {
		fmt.Print(y, " not a power of 2")
	}
}
