package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := [][]int{{1, 2}, {2, 3}, {4, 5}}
	nums2 := [][]int{{1, 4}, {3, 2}, {4, 1}}

	result := mergeArrays(nums1, nums2)
	fmt.Println(result)
}

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
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
