// Problem 2: Even Fibonacci numbers

// Each new term in the Fibonacci sequence is generated 
// by adding the previous two terms. By starting with 1 
// and 2, the first 10 terms will be:
// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...


//////////////////////////////////////////////////////////////////
// By considering the terms in the Fibonacci sequence whose /////
// values do not exceed four million, find the sum of the //////
// even-valued terms. /////////////////////////////////////////
//////////////////////////////////////////////////////////////


// PSUEDOCODE
// Step 1: term1 = 1 
// Step 2: PRINT term1
// Step 3: term2 = 2
// Step 4: PRINT term2
// Step 5: currentTerm = 0
// Step 6: sumOfEvenTerms = 2
// Step 7: WHILE currentTerm <= 4000000
// Step 8: currentTerms = term1 + term2
// Step 9: term1 = term2
// Step 10: term2 = currentTerm
// Step 11: PRINT currentTerm
// Step 12: IF currentTerm is divisible by 2 
// Step 13: sumOfEvenTerms += currentTerm
// Step 14: END IF
// Step 15: END WHILE
// Step 16: PRINT sum


package main

import "fmt"

func main() {
	term1 := uint(1)
	term2 := uint(2)
	currentTerm := uint(0)
	sumOfEvenTerms := uint(2)

	// find the sum of evened-valued term below 4000000
	for ; currentTerm <= 4000000; {
		currentTerm = term1 + term2

		term1 = term2
		term2 = currentTerm

		if currentTerm%2 == 0 {
			sumOfEvenTerms += currentTerm
		}
	}

	// Print sumOfEvenTerms
	fmt.Println("Sum Of Even Terms: ",  sumOfEvenTerms)
}
