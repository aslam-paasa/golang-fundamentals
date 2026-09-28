package main

import "fmt"

func main() {

	/**
	 * Variable Declaration in Go
	 */

	/**
	 * 1. Explicit Type
	 *    - Declare the variable with its type.
	 *    - Syntax: var name type = value
	 */
	var name string = "golang"

	/**
	 * 2. Type Inference
	 *    - Go automatically detects the variable's type from the assigned value.
	 */
	var name2 = "golang"
	var isAdult = true
	var age = 30

	/**
	 * 3. Shorthand Declaration
	 *    - The := syntax declares and initializes a variable.
	 *    - Go automatically detects the type.
	 *    - Can only be used inside functions.
	 */
	name3 := "golang"

	/**
	 * 4. Declare First, Assign Later
	 *    - A variable can be declared without a value.
	 *    - It gets the zero value of its type.
	 */
	var name4 string
	name4 = "golang"

	/**
	 * 5. Float Variables
	 *    - Go can infer the float type automatically.
	 *    - Decimal values are inferred as float64.
	 */
	var price float32 = 50.5
	var price2 = 50.5
	price3 := 50.5

	fmt.Println(name, name2, isAdult, age)
	fmt.Println(name3, name4)
	fmt.Println(price, price2, price3)
}