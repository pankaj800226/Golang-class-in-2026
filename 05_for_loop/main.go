package main

import "fmt"

func main() {

	// while loop
	fmt.Println("While Lopp")
	n:=1
	for n <=3 {
		fmt.Println(n)
		n = n +1
	}

	// for loop
	fmt.Println("for Lopp")
	for i := 1; i <= 10; i++ {

		if i == 3 {
			continue
		}
		
		fmt.Println(i)

	}

	// range 
	fmt.Println("Range ")
	for j := range 11 {
		fmt.Println(j)
	}

}
