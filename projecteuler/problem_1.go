// Problem 1: Multiples of 3 or 5

// If we list all the natural numbers below 10 
// that are multiples of 3 or 5, we get 3, 5, 6 and 9. 
// The sum of these multiples is 23.

///////////////////////////////////////////////////////////////
// Find the sum of all the multiples of 3 or 5 below 1000./////
///////////////////////////////////////////////////////////////


// SOLUTION
// step 1: Count 1 to 999
// step 2: IF number divisible by 3 or 5, add number to total
// step 3: PRINT total

package main

import "fmt"

func main() {

	var counter uint
	total := uint(0)

	for counter=0; counter<1000; counter++ {
		if counter%3 == 0 {
			total += counter
		} else if counter%5 == 0 {
			total += counter
		}
	}
	fmt.Println("Sum: ", total)

}
