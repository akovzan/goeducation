package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	if n > 0 {
		for i := 1; i <= n; i++ {
			fmt.Print(i)
		}
	}
}
