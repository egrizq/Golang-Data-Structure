package loop

import "sort"

// leetCode 1051. Hight Checker

/*

You are given an integer array heights representing the current order that the students are standing in.
Each heights[i] is the height of the ith student in line (0-indexed).

Return the number of indices where heights[i] != expected[i].

Example 1:

Input: heights = [1,1,4,2,1,3]
Output: 3
Explanation:
heights:  [1,1,4,2,1,3]
expected: [1,1,1,2,3,4]
Indices 2, 4, and 5 do not match.

*/

func HeightChecker(heights []int) int {
	copyHeights := make([]int, len(heights))
	copy(copyHeights, heights)
	sort.Ints(copyHeights)

	counter := 0
	for i := range heights {
		if heights[i] != copyHeights[i] {
			counter++
		}
	}

	return counter
}
