package main

import "fmt"

func main() {

	/**
	 * Arrays in Go
	 * - An array is a numbered sequence of elements.
	 * - The size of an array is fixed when it is declared.
	 */

	/**
	 * 1. Zero Values
	 *    - When an array is declared without values, its elements get their zero values.
	 *      - int    -> 0
	 *      - string -> ""
	 *      - bool   -> false
	 */
	var nums [4]int

	fmt.Println(nums)

	/**
	 * 2. Access and Update Array Elements
	 *    - Array indexing starts from 0.
	 *    - Use the index to access or update an element.
	 */
	nums[0] = 1

	fmt.Println(nums[0])
	fmt.Println(nums)

	/**
	 * 3. Array Length
	 *    - len() returns the number of elements in an array.
	 */
	fmt.Println(len(nums))

	/**
	 * 4. Arrays with Different Data Types
	 *    - Arrays can contain elements of different basic types.
	 *    - All elements must have the same type.
	 */
	var vals [4]bool
	vals[2] = true

	fmt.Println(vals)

	var name [3]string
	name[0] = "golang"

	fmt.Println(name)

	/**
	 * 5. Array Initialization
	 *    - An array can be declared and initialized in one line.
	 */
	nums2 := [3]int{1, 2, 3}

	fmt.Println(nums2)

	/**
	 * 6. Two-Dimensional Arrays
	 *    - An array can contain other arrays.
	 *    - This is commonly used to represent a matrix.
	 */
	matrix := [2][2]int{
		{3, 4},
		{5, 6},
	}

	fmt.Println(matrix)

	/**
	 * 7. Advantages of Arrays
	 *    - Fixed size and predictable memory usage.
	 *    - Efficient memory usage.
	 *    - Constant-time access using an index.
	 */
}