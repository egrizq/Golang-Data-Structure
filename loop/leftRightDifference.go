package loop

// leetCode No. 2574 Left and Right Sum Differences

/*

Given a 0-indexed integer array nums, find a 0-indexed integer array answer where:

answer.length == nums.length.
answer[i] = |leftSum[i] - rightSum[i]|.
Where:

leftSum[i] is the sum of elements to the left of the index i in the array nums.
f there is no such element, leftSum[i] = 0.
rightSum[i] is the sum of elements to the right of the index i in the array nums.
If there is no such element, rightSum[i] = 0.
Return the array answer.

Example 1:

Input: nums = [10,4,8,3]
Output: [15,1,11,22]
Explanation: The array leftSum is [0,10,14,22] and the array rightSum is [15,11,3,0].
The array answer is [|0 - 15|,|10 - 11|,|14 - 3|,|22 - 0|] = [15,1,11,22].

*/

func LeftRightDifference(nums []int) []int {
	length := len(nums)

	leftSum, counter := []int{}, 0
	for i := range length {
		leftSum = append(leftSum, counter)
		counter += nums[i]
	}

	counter = 0
	for i := length - 1; i >= 1; i-- {
		counter += nums[i]
	}

	rightSum := []int{}
	for i := 1; i < length; i++ {
		rightSum = append(rightSum, counter)
		counter -= nums[i]
	}
	rightSum = append(rightSum, 0)

	ans := []int{}
	for i := range length {
		sum := leftSum[i] - rightSum[i]
		if sum < 0 {
			sum = -sum
		}
		ans = append(ans, sum)
	}

	return ans
}
