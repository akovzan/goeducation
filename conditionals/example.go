package main

import (
	"fmt"
)

func main() {
	// Conditionals
	a := 1
	b := 2

	if a > b { // 'a' is greater than 'b'
		fmt.Println("Not printed")
	} else if b == a { // 'a' is equal 'b'
		fmt.Println("Not printed")
	} else {
		fmt.Println("Is printed")
	}

	c := 3

	if c > a && c > b { // 'c' greater than 'a' AND 'c' greater than 'b'
		fmt.Println("Is printed")
	}

	if a > b || a > c { // 'a' greater than 'b' OR 'a' greater than 'c'
		fmt.Println("Not printed")
	}

	if a != b { // 'a' is not equal 'b'
		fmt.Println("Is printed")
	}

	/*
		       1 = true, 0 = false


				A   B       A && B      A || B       !A
				________________________________________

				1   1          1           1          0
				1   0          0           1          0
				0   1          0           1          1
				0   0          0           0          1

	*/

	var password string
	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	if password == "1234" { // You can also compare strings
		fmt.Println("Welcome to the app")
	} else {
		fmt.Println("Password is incorrect")
	}

	// Switch
	name := "Denis"

	switch name {
	case "Andrey":
		fmt.Println("Hello, Father!")
	case "Elena":
		fmt.Println("Hello, Mother!")
	case "Denis":
		fmt.Println("Hello, Son!")
	default:
		fmt.Println("I do not know you")
	}

	/*
		This block of code is the same as this:

		if name == "Andrey" {
			fmt.Println("Hello, Father!")
		} else if name == "Elena" {
			fmt.Println("Hello, Mother!")
		} else if name == "Denis" {
			fmt.Println("Hello, Son!")
		} else {
			fmt.Println("I do not know you")
		}

		But in the situations like this 'switch' is better because it much more readable.
	*/
}
