// Problem 3: Largest prime factor

//The prime factors of 13195 are 5, 7, 13 and 29.

////////////////////////////////////////////////////////////////////
// What is the largest prime factor of the number 600851475143 ? // 
////////////////////////////////////////////////////////////////// 

// PSEUDOCODE
// Step 1: theNumber = 600851475143
// Step 2: primeNumber = 2 
// Step 3: WHILE theNumber > 1
// Step 4: IF theNumber % primeNumber == 0
		// theNumber /= primeNumber
		// primeFactorsArr = append(primeNumber) + ' '
		// primeNumber = 2
	// ELSE primeNumber = NextPrimeNumber()
// Step 5: END IF 
// Step 6: END WHILE
// Step 7: Print Max(primeFactorsArr)


// Version 3
// Append prime factor into an uint array - v2
// Changed uint variables to uint64, because constant 600851475143 overflows uint - v3

package math

// Finds the next prime number after a given number
func NextPrimeNumber(currentPN uint64) uint64 {
	currentPN++
	for i:= uint64(2); i <= currentPN; i++ {
		if currentPN % i == 0 && currentPN != i {
			return NextPrimeNumber(currentPN)
		} else if currentPN == i {
			return currentPN
		}
	}
	return currentPN
}


// Finds the largest number within an array
func Max(arr []uint64) uint64 {
	maximum := uint64(0)
	for _, value := range arr {
		if value > maximum {
			maximum = value
		}
	}
	return maximum
}


// Find the prime factors of a number
func PrimeFactors(theNumber uint64) []uint64 {
	primeNumber := uint64(2)
	primeFactorsArr := []uint64{}

	for theNumber > 1 {
		if theNumber % primeNumber == 0 {
			theNumber /= primeNumber
			primeFactorsArr = append(primeFactorsArr, primeNumber)
			primeNumber = 2
		} else {
			primeNumber = NextPrimeNumber(primeNumber)
		}
	}
	return primeFactorsArr

}


// Finds the largest prime factor of a number
func LargestPrimeFactor(theNumber uint64) uint64 {
	return Max(PrimeFactors(theNumber))
}

