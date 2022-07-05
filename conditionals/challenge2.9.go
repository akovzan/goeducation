package main

import (
	"fmt"
)

func main() {
	var x1 int
	var x2 int
	var y1 int
	var y2 int
	fmt.Print("x1 = ")
	fmt.Scan(&x1)
	fmt.Print("x2 = ")
	fmt.Scan(&x2)
	fmt.Print("y1 = ")
	fmt.Scan(&y1)
	fmt.Print("y2 = ")
	fmt.Scan(&y2)

	if (x1+x2+y1+y2)%2 == 0 {
		fmt.Println("клетки одного цвета")
	} else {
		fmt.Println("клетки разного цвета")
	}
}
