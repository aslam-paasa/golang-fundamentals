package main

import (
	"fmt"
	"maps"
)

func main() {

	/**
	 * Maps in Go
	 * - A map stores data as key-value pairs.
	 * - Similar to a hash map, object, or dictionary.
	 * - Each key must be unique.
	 */

	/**
	 * 1. Create a Map
	 *    - make() can be used to create an empty map.
	 *    - The first type is the key type.
	 *    - The second type is the value type.
	 */
	m := make(map[string]string)

	/**
	 * 2. Set Elements
	 *    - Use map[key] = value to add or update an element.
	 */
	m["name"] = "golang"
	m["area"] = "backend"

	/**
	 * 3. Get Elements
	 *    - Use map[key] to access a value.
	 *    - If the key does not exist, Go returns the zero value of the value's type.
	 */
	fmt.Println(m["name"], m["area"])

	/**
	 * 4. Map with Integer Values
	 *    - The key and value can have different types.
	 *    - Here, the keys are strings and values are integers.
	 */
	m2 := make(map[string]int)

	m2["age"] = 30
	m2["price"] = 50

	fmt.Println(m2["phone"])
	fmt.Println(len(m2))

	/**
	 * 5. Delete an Element
	 *    - delete() removes an element using its key.
	 */
	delete(m2, "price")

	fmt.Println(m2)

	/**
	 * 6. Clear a Map
	 *    - clear() removes all elements from the map.
	 */
	clear(m2)

	fmt.Println(m2)

	/**
	 * 7. Map Literal
	 *    - A map can be created and initialized in one line.
	 */
	m3 := map[string]int{
		"price":  40,
		"phones": 3,
	}

	fmt.Println(m3)

	/**
	 * 8. Check Whether a Key Exists
	 *    - The second value returned by a map lookup tells whether the key exists.
	 *      a. value -> value stored for the key.
	 *      b. ok    -> true if the key exists, otherwise false.
	 */
	v, ok := m3["phones"]

	fmt.Println(v)

	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	/**
	 * 9. Compare Maps
	 *    - maps.Equal() checks whether two maps contain the same key-value pairs.
	 */
	m4 := map[string]int{
		"price":  40,
		"phones": 3,
	}

	m5 := map[string]int{
		"price":  40,
		"phones": 8,
	}

	fmt.Println(maps.Equal(m4, m5))
}