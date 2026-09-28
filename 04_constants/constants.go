package main

import "fmt"

func main() {

	/**
	 * Constants in Go
	 * - A constant is a value that cannot be changed after it is declared.
	 */

	/**
	 * 1. Declare a Constant
	 *    - Use the const keyword to declare a constant.
	 *    - Constants can be declared without :=.
	 */
	const name = "golang"
	const age = 30

	fmt.Println(name, age)

	/**
	 * 2. Multiple Constants
	 *    - Multiple constants can be grouped together using parentheses.
	 */
	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}
