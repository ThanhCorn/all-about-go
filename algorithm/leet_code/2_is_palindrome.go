package leetcode

// Given an integer x, return true if x is a
// palindrome, and false otherwise.

func isPalindrome(x int) bool {
	// Negative numbers cannot be palindromes
	if x < 0 {
		return false
	}

	original := x
	reversed := 0

	// Reverse the number
	for x != 0 {
		lastDigit := x % 10
		reversed = reversed*10 + lastDigit
		x /= 10
	}

	// Check if the original number is equal to the reversed number
	return original == reversed
}
