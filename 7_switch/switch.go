package main

import (
	"fmt"
	"time"
)

func main() {

	/**
	 * Switch Statement in Go
	 * - switch is used to check a value against multiple cases.
	 * - Go automatically stops after a matching case.
	 * - break is not required.
	 */

	/**
	 * 1. Simple Switch
	 *    - Checks a value against different cases.
	 *    - default runs when no case matches.
	 */
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

	/**
	 * 2. Multiple Values in a Case
	 *    - Multiple values can be checked in a single case.
	 */
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend")
	default:
		fmt.Println("it's workday")
	}

	/**
	 * 3. Type Switch
	 *    - Used to check the type of a value.
	 *    - The type is checked using i.(type).
	 */
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("it's an integer")
		case string:
			fmt.Println("it's a string")
		case bool:
			fmt.Println("it's a boolean")
		default:
			fmt.Println("other")
		}
	}

	whoAmI(55)
}