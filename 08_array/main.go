package main

import "fmt"

func main() {
	fmt.Println("Hello array")

	// int > 0, string > "", boll > false

	// two declear single line array (int)
	array := [4]int{1, 2, 3, 4}
	fmt.Println(array)

	// one by one
	var nums [4]int

	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums[3] = 4

	fmt.Println(nums)

	// string
	var name [4]string

	name[0] = "golang"
	fmt.Println(name)

	// boll

	var bol [4]bool

	bol[0] = false

	fmt.Println(bol)

	// 2d array
	numbers := [2][2]int{{1, 2}, {3, 4}}

	fmt.Println(numbers)

}
