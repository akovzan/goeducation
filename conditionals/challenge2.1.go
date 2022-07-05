package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	if a > b {
		fmt.Println("greater")
	} else if a < b {
		fmt.Println("lower")
	} else {
		fmt.Println("equal")
	}
}
