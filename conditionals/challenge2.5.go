package main

import (
	"fmt"
)

func main() {
	var usename string
	var password string
	var a int
	var b int
	fmt.Println("Enter usename: ")
	fmt.Scan(&usename)
	fmt.Println("Enter password: ")
	fmt.Scan(&password)

	if usename == "admin" && password == "1234" {
		fmt.Println("success massage")
		fmt.Println("enter a ")
		fmt.Scan(&a)
		fmt.Println("enter b ")
		fmt.Scan(&b)
		fmt.Println("sum a + B = ", a+b)
	} else {
		fmt.Println("credentials are invalid")
	}
}
