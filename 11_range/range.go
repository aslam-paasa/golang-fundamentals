package main

import "fmt"

func main() {

	/**
	 * Range in Go
	 * - range is used to iterate over data structures such as slices, maps, strings, 
	 *   and arrays.
	 */

	/**
	 * 1. Range Over a Slice
	 *    - Returns the index and value of each element.
	 *      a. index -> position of the element.
	 *      b. value -> element at that position.
	 */
	nums := []int{6, 7, 8}

	for index, value := range nums {
		fmt.Println("index:", index, " - ", "value:", value)
	}

	/**
	 * 2. Range Over a Map
	 *    - Returns the key and value of each element.
	 */
	m := map[string]string{
		"fname": "john",
		"lname": "doe",
	}

	for key, value := range m {
		fmt.Println("key:", key, "value:", value)
	}

	/**
	 * 3. Range Over Map Keys
	 *    - If only one variable is used, range returns the key.
	 */
	for key := range m {
		fmt.Println("key:", key)
	}

	/**
	 * 4. Range Over a String
	 *    - range iterates over a string by Unicode code points (runes).
	 *    - The first value is the starting byte index of the rune.
	 *    - The second value is the rune itself.
	 */
	for i, c := range "golang" {
		fmt.Println("index:", i, "-", "char:", string(c))
	}
}