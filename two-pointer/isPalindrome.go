package twopointer

import (
	"strings"
	"unicode"
)

// leetCode 125. Valid Palindrome

/*

A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and
removing all non-alphanumeric characters, it reads the same forward and backward.
Alphanumeric characters include letters and numbers.
Given a string s, return true if it is a palindrome, or false otherwise.

Example 1:
Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

Example 2:
Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.
*/

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)

	l, r := 0, len(s)-1
	for l < r {
		leftChar := s[l]
		rightChar := s[r]

		if !unicode.IsNumber(rune(leftChar)) && !unicode.IsLetter(rune(leftChar)) {
			l++
		} else if !unicode.IsNumber(rune(rightChar)) && !unicode.IsLetter(rune(rightChar)) {
			r--
		} else if leftChar == rightChar {
			l++
			r--
		} else {
			return false
		}
	}

	return true
}
