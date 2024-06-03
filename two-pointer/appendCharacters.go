package twopointer

// leetCode 2486. Append Characters to String to Make Subsequence

/*
You are given two strings s and t consisting of only lowercase English letters.
Return the minimum number of characters that need to be appended to the end of s so that t becomes a subsequence of s.

Example 1:
Input: s = "coaching", t = "coding"
Output: 4
Explanation: Append the characters "ding" to the end of s so that s = "coachingding".
Now, t is a subsequence of s ("coachingding").
It can be shown that appending any 3 characters to the end of s will never make t a subsequence.

Example 2:
Input: s = "abcde", t = "a"
Output: 0
Explanation: t is already a subsequence of s ("abcde").

*/

func AppendCharacters(s string, t string) int {
	var l, r, ans int

	for {
		if s[l] == t[r] {
			r++
			l++
		} else {
			l++
		}

		if l == len(s) {
			ans += len(t[r:])
			return ans
		}

		if r == len(t) {
			return ans
		}
	}
}
