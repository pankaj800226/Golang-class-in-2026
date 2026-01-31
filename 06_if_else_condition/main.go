package main

import "fmt"

func main() {

	// if else

	var age = 20

	if age >= 20 {
		fmt.Println("you can drive")
	} else if age >= 16 {
		fmt.Println("you age 16 can not drive")
	} else if age >= 12 {
		fmt.Println("you age 12 can not drive")

	} else {
		fmt.Println("you can not drive")

	}

	// OR operator and and && operator

	role := "admin"
	hasParmistion := false

	// or

	// if role  == "admin" || hasParmistion {
	// 	fmt.Println("yes")
	// }

	// and 

	if role == "admin" && hasParmistion {
			fmt.Println("yes")
		}
	}

