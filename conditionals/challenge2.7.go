package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var c int
	fmt.Println("Enter a: ")
	fmt.Scan(&a)
	fmt.Println("Enter b: ")
	fmt.Scan(&b)
	fmt.Println("Enter c: ")
	fmt.Scan(&c)
	if a == b && b == c {
		fmt.Println("All of them are the same")
	} else if a == b || b == c || a == c {
		fmt.Println("Two of them are the same")
	} else {
		fmt.Println("All of them are different")
	}

}
