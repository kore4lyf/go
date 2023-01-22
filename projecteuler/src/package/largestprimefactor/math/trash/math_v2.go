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
		// primeFactorsStr = string(primeNumber) + ' '
		// primeNumber = 2
	// ELSE primeNumber = NextPrimeNumber()
// Step 5: END IF 
// Step 6: END WHILE
// Step 7: primeFactorsArr = strings.Split(primeFactorsStr, " ")
// Step 8: Print Max(primeFactorsArr)


// Version 2
// Append prime factor into an uint array - v2

package math

// Finds the next prime number after a given number
func NextPrimeNumber(currentPN uint) uint {
	currentPN++
	for i:= uint(2); i <= currentPN; i++ {
		if currentPN % i == 0 && currentPN != i {
			return NextPrimeNumber(currentPN)
		} else if currentPN == i {
			return currentPN
		}
	}
	return currentPN
}


// Finds the largest number within an array
func Max(arr []uint) uint {
	maximum := uint(0)
	for _, value := range arr {
		if value > maximum {
			maximum = value
		}
	}
	return maximum
}


// Finds the largest prime factor of a number
func LargestPrimeFactor(theNumber uint) uint {
	primeNumber := uint(2)
	primeFactorsArr := []uint{}

	for theNumber > 1 {
		if theNumber % primeNumber == 0 {
			theNumber /= primeNumber
			primeFactorsArr = append(primeFactorsArr, primeNumber)
			primeNumber = 2
		} else {
			primeNumber = NextPrimeNumber(primeNumber)
		}
	}

	return Max(primeFactorsArr)
}

