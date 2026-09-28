package main

import "fmt"

func main() {

	/**
	 * Loops in Go:
	 * - Go has only one looping construct: for
	 * - The for loop can be used like a while loop, infinite loop, or traditional 
	 *   for-loop.
	 */

	/**
	 * 1. While-Style Loop
	 *    - Go does not have a separate while keyword.
	 *    - Use for with only a condition.
	 */
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	/**
	 * 2. Infinite Loop
	 *    - A for loop without a condition runs forever.
	 *    - Use break when you want to stop the loop.
	 */
	for {
		fmt.Println("Running...")
		break
	}

	/**
	 * 3. Classic For Loop
	 *    - Contains three parts:
	 *      1. Initialization
	 *      2. Condition
	 *      3. Increment
	 */
	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}

	/**
	 * 4. break
	 *    - Stops the loop immediately.
	 */
	for i := 0; i <= 3; i++ {
		if i == 2 {
			break
		}

		fmt.Println(i)
	}

	/**
	 * 5. continue
	 *    - Skips the current iteration and moves to the next iteration.
	 */
	for i := 0; i <= 3; i++ {
		if i == 2 {
			continue
		}

		fmt.Println(i)
	}

	/**
	 * 6. Range Over Integers
	 *    - Go 1.22+ supports ranging directly over integers.
	 *    - range 11 produces values from 0 to 10.
	 */
	for i := range 11 {
		fmt.Println(i)
	}
}