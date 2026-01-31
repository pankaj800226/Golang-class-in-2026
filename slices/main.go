package main

import "fmt"

func main() {
	// you nit fixed array

	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	var nums = make([]int, 2, 5)

	// nums := []int{}
	nums[0] = 1
	nums[1] = 2
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)
	nums = append(nums, 6)
	nums = append(nums, 7)
	nums = append(nums, 8)
	nums = append(nums, 9)
	nums = append(nums, 10)

	fmt.Println(nums)
	// capacity
	fmt.Println(cap(nums))

}
