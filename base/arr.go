package main

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

//

// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Output: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:

// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]

import (
	"fmt"
)

func printArr(vals []interface{}) {
	for idx, value := range vals {
		fmt.Printf("%d: %s\n", idx, value)
	}
}

func main() {
	// 数组
	fmt.Println("Example 1:")
	s := []string{"a", "b", "c"}
	fmt.Println("capacity:", cap(s))
	s_interface := make([]interface{}, 3)
	for i, v := range s {
		s_interface[i] = v
	}
	printArr(s_interface)
	return
}
