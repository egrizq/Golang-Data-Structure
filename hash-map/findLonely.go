package hashmap

// leetCode No. 2150 Find All Lonely Numbers in the Array

/*
You are given an integer array nums. A number x is lonely when it appears only once,
and no adjacent numbers (i.e. x + 1 and x - 1) appear in the array.

Return all lonely numbers in nums. You may return the answer in any order.

Example 1:
Input: nums = [10,6,5,8]
Output: [10,8]
Explanation:
- 10 is a lonely number since it appears exactly once and 9 and 11 does not appear in nums.
- 8 is a lonely number since it appears exactly once and 7 and 9 does not appear in nums.
- 5 is not a lonely number since 6 appears in nums and vice versa.
Hence, the lonely numbers in nums are [10, 8].
Note that [8, 10] may also be returned.

*/

func FindLonely(nums []int) []int {
	hashMap := make(map[int]int)

	for _, num := range nums {
		hashMap[num]++
	}

	ans := []int{}
	for key, value := range hashMap {
		if value == 1 && hashMap[key-1] == 0 && hashMap[key+1] == 0 {
			ans = append(ans, key)
		}
	}

	return ans
}
