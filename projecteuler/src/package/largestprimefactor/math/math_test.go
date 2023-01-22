package math

import "testing"


func TestLargestPrimeFactor(t *testing.T) {
	num := LargestPrimeFactor(13195)

	if num != 29 {
		t.Error(
				"For Number 13195",
				"\n Expected: ", 29,
				"\n Got: ", num)
	}

}
