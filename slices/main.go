package main

import (
	"fmt"
	"slices"
)

func main() {
	// you nit fixed array

	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// var nums = make([]int, 2, 5)

	// nums := []int{}
	// nums[0] = 1
	// nums[1] = 2
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// nums = append(nums, 5)
	// nums = append(nums, 6)
	// nums = append(nums, 7)
	// nums = append(nums, 8)
	// nums = append(nums, 9)
	// nums = append(nums, 10)

	// fmt.Println(nums)
	// capacity
	// fmt.Println(cap(nums))

	// copy function
	// var nums = make([]int, 0, 5)

	// nums = append(nums, 2)

	// var num2 = make([]int,len(nums))

	// copy(num2,nums)
	// fmt.Println(nums,num2)

	// slice operator
	// var nums = []int{1, 2, 3, 4}

	// fmt.Println(nums[0:3])
	// fmt.Println(nums[:2])
	// fmt.Println(nums[2:])

	// slice package check equal or not

	var num1 = []int{1, 2, 3}
	var num2 = []int{1, 2, 3}

	fmt.Println(slices.Equal(num1, num2))
	
}