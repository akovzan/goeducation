package main

import (
	"fmt"
)

func main() {
	var a int
	var x int
	var y int
	var b = 20
	var c = 5
	var d = 1
	fmt.Print("emount = ")
	fmt.Scan(&a)
	if a > b {
		fmt.Println(a/b, "of 20$")
	}
	x = a % b
	if x > c {
		fmt.Println(x/c, "of 5$")
	}
	y = x % c
	if y > 0 {
		fmt.Println(y/d, "of 1$")
	}

}
