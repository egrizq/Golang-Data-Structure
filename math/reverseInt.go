package math

import "math"

// leetCode 7. Reverse Int

/*

Given a signed 32-bit integer x, return x with its digits reversed.
If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

Example 1:
Input: x = 123
Output: 321

Example 2:
Input: x = -123
Output: -321

*/

func Reverse(x int) int {
	reverseNumber := 0

	for x != 0 {
		lastNumber := x % 10
		x /= 10
		reverseNumber = reverseNumber*10 + lastNumber

		if reverseNumber > math.MaxInt32 || reverseNumber < math.MinInt32 {
			return 0
		}
	}

	return reverseNumber
}
