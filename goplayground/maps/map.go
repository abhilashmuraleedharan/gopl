/*
 * Map is a built-in datastructure that associate values of one type (the key)
 * with values of another type (the element or the value).
 * The entries will not be stored in order like in arrays.

 * The key can be of any type for which the equality operator is defined.
 * For e.g. you can use a string type as a key, but you cannot use a slice as a key.
 *
 * Like slices, maps hold references to an underlying data structure. So if you pass a map
 * to a function that changes the content of that map, the changes will be visible in the caller.
 */
package main

import "fmt"

func main() {
	// creating a map [EmployeeName: EmployeeID] using composite literal
	m := map[string]int{
		"Jerremy": 26467,
		"James":   31567,
		"Richard": 67812,
	}

	fmt.Println(m)

	// Accessing a value using the key
	fmt.Println("Employee Id of James:", m["James"])

	// Trying to access a value using a non-existent key
	fmt.Println("Employee Id of Abby: ", m["Abby"])

	/*
	 * Since no errors will be thrown when a non-existent key is used,
	 * to help us know whether a key:value is present in map, when we access a key,
	 * along with value it returns an extra boolean value which will be false if key is non-existant.
	 * There is a Comma Idiom method to check this boolean value as shown below
	 */
	if v, ok := m["Abby"]; ok { // This is the Comma Ok idiom
		fmt.Println("Employee Id of Abby: ", v)
	} else {
		fmt.Println("There is no employee named Abby in the database")
	}

	// Add a new key:value pair to the map
	m["Abby"] = 34512

	// Using for loop with range clause, we can loop over a map
	for k, v := range m {
		fmt.Println(k, v)
	}

	/*
	 * To delete a key:value pair from map, use built-in function delete
	 * Since deleting a non-existant key won't throw any error, we can use
	 * the Comma Idiom method to confirm a key exists before deleting it
	 */
	if _, ok := m["Abby"]; ok {
		delete(m, "Abby") // delete(<map name>, <key to delete>)
	}

	fmt.Println(m)
}
