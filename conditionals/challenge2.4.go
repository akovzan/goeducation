package main

import (
	"fmt"
)

func main() {
	var usename string
	var password string
	fmt.Println("Enter usename: ")
	fmt.Scan(&usename)
	fmt.Println("Enter password: ")
	fmt.Scan(&password)

	if usename == "admin" && password == "1234" {
		fmt.Println("success massage")
	} else {
		fmt.Println("credentials are invalid")
	}
}
