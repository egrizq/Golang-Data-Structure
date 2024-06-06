package loop

// leetCode 1221. Split a String in Balanced Strings

/*

Balanced strings are those that have an equal quantity of 'L' and 'R' characters.
Given a balanced string s, split it into some number of substrings such that:

- Each substring is balanced.
- Return the maximum number of balanced strings you can obtain.

Example 1:

Input: s = "RLRRLLRLRL"
Output: 4
Explanation: s can be split into "RL", "RRLL", "RL", "RL", each substring contains same number of 'L' and 'R'.

*/

func BalancedStringSplit(s string) int {
	ans := 0
	countR, countL := 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			countR++
		} else {
			countL++
		}

		if countR == countL {
			ans++
			countR, countL = 0, 0
		}
	}

	return ans
}
