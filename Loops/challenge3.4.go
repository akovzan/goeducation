package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	if n > 2 {
		for i := 2; i <= n; i++ {
			if i%2 != 0 {
				continue
			}

			fmt.Print(i)
		}
	}
}
