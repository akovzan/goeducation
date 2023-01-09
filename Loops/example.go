package main

import (
	"fmt"
)

func main() {
	// Loops

	// Three-component loop
	for i := 1; i <= 5; i++ { 
		/*

	 		First component is 'i := 1'. This operator runs one time before the loop.

	 		Second component is 'i <= 5'. This is the condition that indicates whether
	 		the loop should continue or not. Thick check runs in the end of every iteration.

	 		This component is 'i++'. This operator runs in the end of every iteration but before
	 		the check.
	 	*/
		fmt.Printf("%v ", i)  // 1 2 3 4 5
	}
	fmt.Println()


	// All components of the loop are optional (not required) and could be skipped.

	n := 3

	for ; n >= 1; n-- { // Skipping first component
		fmt.Printf("%v ", n) // 3 2 1
	}
 	fmt.Println()


 	n = 2

 	for n <= 9 { // First and third components are skipped.
 		fmt.Printf("%v ", n) // 2 4 6 8
 		n += 2
 	}
 	fmt.Println()


 	n = 1
 	for { // All components is skipped. This is called an 'Infinite loop'
 		n++
 		if n % 2 == 0 {
 			continue // This operator finishes current iteration and goes to the next iteration.
 		}

 		fmt.Printf("%v ", n)

 		if n >= 10 {
 			fmt.Println()
 			break // This operator finishes the loop.
 		}
 	}

 	// 'continue' and 'breaks' operators can be used in any kind of loop
 	for index := 1; index <= 100; index++ {
 		if index % 3 != 0 {
 			continue
 		}

 		if index == 15 {
 			break
 		}

 		fmt.Printf("%v ", index)
 	}
 	fmt.Println()
}
