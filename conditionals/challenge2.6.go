package main

import (
	"fmt"
)

func main() {
	var usename string
	var password string
	var a int
	var b int
	var operator string
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
		fmt.Scan(&operator)
		if operator == "+" {
			fmt.Println("a + b = ", a+b)
		} else if operator == "-" {
			fmt.Println("a - b = ", a-b)
		} else if operator == "*" {
			fmt.Println("a * b = ", a*b)
		} else if operator == "/" {
			fmt.Println(a, " / ", b, " = ", a/b)
		} else {
			fmt.Println("error")
		}
	} else {
		fmt.Println("credentials are invalid")
	}
}
