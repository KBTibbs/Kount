package happynumber

import "fmt"

/*****************************************************************************/
/* Package happynumber contains utility functions for working with happy     */
/* numbers. Happy Numbers are any positive integer that when the digits are  */
/* split, squared and summed repeatedly until they result in a single digit  */
/* value.                                                                    */
/* Numbers that result in a 1 are Happy Numbers.                             */
/* Numbers that result in a 4 are not HAppy Numbers.                         */
/*****************************************************************************/

func TestNumber(num int) (bool, error) {
	// It is overkill to use a switch for this instead of a normal If-then-else, but I was
	// rather taken with the the idea of a swtich without a conditional in the documentation
	// and so I used it.
	switch {
	case num < 1:
		return false, fmt.Errorf("TestNumber: cannot test numbers below 1: %d", num)
	case num == 1:
		return true, nil
	case num == 4: 
	    // testing for 16, 37, 58, 89, 145, 42, 20 might be an optimization
	    // All unhappy numbers will eventually fall into this looping pattern
		return false, nil
	default:
		return TestNumber(sumSquareOfDigits(num))
	}
}

func sumSquareOfDigits(x int) int {
	var digit, sum int

	for x > 0 {
		// Modulus pops off the smallest ordinal, which is technically backwards
		// but since addition is commutative, it doesn't matter.
		digit = x % 10
		sum = sum + (digit * digit)
		// shift all the numbers down one ordinal to replace the previously removed digit
		x = x / 10
	}
	return sum
}
