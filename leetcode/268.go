package main

import "sort"

// Given an array nums containing n distinct numbers in the range [0, n],
// return the only number in the range that is missing from the array.
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/missing-number
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 最笨的方法
func missingNumber_1(nums []int) int {
	sort.Ints(nums)
	numsLen := len(nums)
	if nums[0] != 0 {
		return 0
	}
	if nums[numsLen-1] != numsLen {
		return numsLen
	}
	for k,v := range nums {
		if k != v {
			return k
		}
	}
	return 0
}

// 利用和相减
func missingNumber_2(nums []int) int {
	var length = len(nums)
	var sum = length * (length + 1) / 2
	for _, v := range nums {
		sum -= v
	}
	return sum
}

// 利用异或做
func missingNumber_3(nums []int) int {
	ret := len(nums)
	for k, v := range nums {
		ret ^= k
		ret ^= v
	}
	return ret
}

func main() {
	println(missingNumber_3([]int{2,1}))
}
