package main

import (
	"fmt"
	"time"
)

func main() {

	// simple switch
	i := 3

	switch i {

	case 1:
		fmt.Println("one")

	case 2:
		fmt.Println("two")

	case 3:
		fmt.Println("three")

	default:
		fmt.Println("other")

	}

	// multiple conditon switch

	switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend")

	default:
		fmt.Println("it's workday")
	}

	// type switch case
	gyan := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("its integer")
		case string:
			fmt.Println("its string")
		case bool:
			fmt.Println("its boolean")
		default:
			fmt.Println("other")
		}
	}

	gyan(90)
}
