// Problem 3: Largest prime factor

//The prime factors of 13195 are 5, 7, 13 and 29.

////////////////////////////////////////////////////////////////////
// What is the largest prime factor of the number 600851475143 ? // 
//////////////////////////////////////////////////////////////////  

// PSEUDOCODE
// Step 1: theNumber = 600851475143
// Step 2: primeNumber = 2 
// Step 3: WHILE primeNumber > 1
// Step 4: IF theNumber % primeNumber == 0
		// theNumber /= primeNumber
		// primeFactorsStr = string(primeNumber) + ' '
		// primeNumber = 2
	// ELSE primeNumber = nextPrimeNumber()
// Step 5: END IF 
// Step 6: END WHILE
// Step 7: primeFactorsArr = strings.Split(primeFactorsStr, " ")
// Step 8: Print LargestNumber


package main

import (
		"fmt"
		"package/largestprimefactor/math"
)

func main() {
	fmt.Println("The Largest Prime Factor of 600851475143 is: ", math.LargestPrimeFactor(600851475143))
}
