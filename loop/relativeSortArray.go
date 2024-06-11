package loop

import "sort"

// leetCode No. 1122 Relative Sort Array

/*

Given two arrays arr1 and arr2, the elements of arr2 are distinct, and all elements in arr2 are also in arr1.

Sort the elements of arr1 such that the relative ordering of items in arr1 are the same as in arr2.
Elements that do not appear in arr2 should be placed at the end of arr1 in ascending order.

Example 1:

Input: arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
Output: [2,2,2,1,4,3,3,9,6,7,19]
Example 2:

Input: arr1 = [28,6,22,8,44,17], arr2 = [22,28,8,6]
Output: [22,28,8,6,17,44]

*/

func RelativeSortArray(arr1 []int, arr2 []int) []int {
	hashMap := make(map[int]int)
	for _, num := range arr1 {
		hashMap[num]++
	}

	ans, pointer := []int{}, 0
	for pointer < len(arr2) {
		if hashMap[arr2[pointer]] > 0 {
			ans = append(ans, arr2[pointer])
			hashMap[arr2[pointer]]--
		} else {
			pointer++
		}
	}

	lastNumber := []int{}
	for key, value := range hashMap {
		if value > 0 {
			lastNumber = append(lastNumber, key)
		}
	}

	pointer, temp := 0, []int{}
	for pointer < len(lastNumber) {
		if hashMap[lastNumber[pointer]] > 0 {
			temp = append(temp, lastNumber[pointer])
			hashMap[lastNumber[pointer]]--
		} else {
			pointer++
		}
	}
	sort.Ints(temp)
	ans = append(ans, temp...)

	return ans
}
