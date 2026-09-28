package main

import (
	"fmt"
	"slices"
)

func main() {

	/**
	 * Slices in Go
	 * - A slice is a dynamic and flexible sequence of elements.
	 * - Unlike arrays, slices do not have a fixed size.
	 * - Slices are one of the most commonly used constructs in Go.
	 */

	/**
	 * 1. Nil Slice
	 *    - A slice declared without initialization is nil.
	 *    - Its length is 0.
	 */
	var nums []int

	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	/**
	 * 2. Create a Slice using make()
	 *    - make() can be used to create a slice with a specified length and capacity.
	 *      a. length   -> number of elements currently in the slice.
	 *      b. capacity -> total space available before the underlying array needs to grow.
	 */
	var nums2 = make([]int, 0, 5)

	fmt.Println(cap(nums2))
	fmt.Println(nums2 == nil)

	/**
	 * 3. Create an Empty Slice
	 *    - A slice can also be initialized using {}.
	 */
	nums3 := []int{}

	/**
	 * 4. append()
	 *    - append() adds elements at the end of the slice.
	 *    - It returns the updated slice.
	 */
	nums3 = append(nums3, 1)
	nums3 = append(nums3, 2)

	nums3[0] = 3
	nums3[1] = 5

	fmt.Println(nums3)
	fmt.Println(cap(nums3))
	fmt.Println(len(nums3))

	/**
	 * 5. copy()
	 *    - copy() copies elements from one slice to another.
	 *    - The destination slice must have enough length
	 *      to receive the elements.
	 */
	nums4 := make([]int, 0, 5)
	nums4 = append(nums4, 2)

	nums5 := make([]int, len(nums4))

	copy(nums5, nums4)

	fmt.Println(nums4, nums5)

	/**
	 * 6. Slice Operator
	 *    - Use [start:end] to create a sub-slice.
	 *    - The start index is included.
	 *    - The end index is excluded.
	 */
	nums6 := []int{1, 2, 3, 4, 5}

	fmt.Println(nums6[0:1])
	fmt.Println(nums6[:2])
	fmt.Println(nums6[1:])

	/**
	 * 7. Compare Slices
	 *    - The slices package provides useful functions for working with slices.
	 *    - slices.Equal() checks whether two slices contain the same elements in the
	 *      same order.
	 */
	nums7 := []int{1, 2, 3}
	nums8 := []int{1, 2, 4}

	fmt.Println(slices.Equal(nums7, nums8))

	/**
	 * 8. Two-Dimensional Slices
	 *    - A slice can contain other slices.
	 *    - This can be used to represent a matrix-like structure.
	 */
	nums9 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(nums9)
}