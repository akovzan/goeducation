package main

import (
	"fmt"
)

func main() {
	var i int
	var n int
	var s int
	fmt.Scan(&i)

	for ; i > 0; i = i / 10 {
		n = i % 10
		s = s + n
	}
	fmt.Print(s)

}
