package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	var b float64
	fmt.Scan(&n)
	b = (math.Sqrt(n))
	//if b%1 = 0 {

	fmt.Println(b)
	fmt.Printf("%v,%t", b, b)
	//}
}
