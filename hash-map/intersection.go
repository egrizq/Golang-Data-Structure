package hashmap

// leetCode 349. Intersection of Two Arrays

/*
Given two integer arrays nums1 and nums2, return an array of their intersection.
Each element in the result must be unique and you may return the result in any order.

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Explanation: [4,9] is also accepted.
*/

func Intersection(nums1 []int, nums2 []int) []int {
	hashMap := make(map[int]bool)

	for _, value := range nums1 {
		hashMap[value] = false
	}

	for _, value := range nums2 {
		if _, ok := hashMap[value]; ok {
			hashMap[value] = true
		}
	}

	ans := []int{}
	for key := range hashMap {
		if hashMap[key] {
			ans = append(ans, key)
		}
	}

	return ans
}
