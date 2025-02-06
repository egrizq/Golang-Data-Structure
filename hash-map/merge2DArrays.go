package hashmap

import "sort"

// Example 1:
//
// Input: nums1 = [[1,2],[2,3],[4,5]], nums2 = [[1,4],[3,2],[4,1]]
// Output: [[1,6],[2,3],[3,2],[4,6]]
// Explanation: The resulting array contains the following:
// - id = 1, the value of this id is 2 + 4 = 6.
// - id = 2, the value of this id is 3.
// - id = 3, the value of this id is 2.
// - id = 4, the value of this id is 5 + 1 = 6.
// Example 2:
//
// Input: nums1 = [[2,4],[3,6],[5,5]], nums2 = [[1,3],[4,3]]
// Output: [[1,3],[2,4],[3,6],[4,3],[5,5]]
// Explanation: There are no common ids, so we just include each id with its value in the resulting list.

func MergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	mapp := make(map[int]int)

	for _, v := range nums1 {
		mapp[v[0]] = v[1]
	}

	for _, v := range nums2 {
		mapp[v[0]] += v[1]
	}

	result := [][]int{}
	for i, v := range mapp {
		result = append(result, []int{i, v})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	return result
}
