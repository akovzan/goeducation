package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var c int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	if a+c > b && a+b > c && c+b > a {
		fmt.Println("can exist")
	} else {
		fmt.Println("can not exist")
	}
	if a*a == b*b+c*c || b*b == a*a+c*c || c*c == b*b+a*a {
		fmt.Println("right triangle")
	}
}
