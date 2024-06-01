package hashmap

// leetCode 136. Single Number

/*

Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.

Example 1:

Input: nums = [2,2,1]
Output: 1

*/

func SingleNumber(nums []int) int {
	hashMap := make(map[int]int)

	for _, number := range nums {
		hashMap[number] += 1
	}

	var answer int
	for key, value := range hashMap {
		if value == 1 {
			answer = key
		}
	}

	return answer
}
