package main

import "fmt"

func main() {

	/**
	 * Conditional Statements in Go
	 * - Go uses if, else if, and else for conditions.
	 */

	/**
	 * 1. if-else
	 *    - Executes one block when the condition is true.
	 *    - Otherwise, the else block is executed.
	 */
	age := 16

	if age >= 18 {
		fmt.Println("person is an adult")
	} else {
		fmt.Println("person is not an adult")
	}

	/**
	 * 2. if-else if-else
	 *    - Used when there are multiple conditions.
	 *    - Conditions are checked from top to bottom.
	 */
	age = 10

	if age >= 18 {
		fmt.Println("person is an adult")
	} else if age >= 12 {
		fmt.Println("person is teenager")
	} else {
		fmt.Println("person is a kid")
	}

	/**
	 * 3. Logical Operators
	 *    - Multiple conditions can be combined.
	 *    - && means AND.
	 *    - || means OR.
	 *    - ! means NOT.
	 */
	role := "admin"
	hasPermissions := false

	if role == "admin" && hasPermissions {
		fmt.Println("yes")
	}

	/**
	 * 4. Variable Declaration Inside if
	 *    - A variable can be declared inside the if statement.
	 *    - The variable is only available inside the if-else block.
	 */
	if age := 20; age >= 18 {
		fmt.Println("person is an adult", age)
	} else if age >= 12 {
		fmt.Println("person is teenager", age)
	}

	/**
	 * 5. Ternary Operator
	 *    - Go does not have a ternary operator.
	 *    - Use a normal if-else statement instead.
	 */
}