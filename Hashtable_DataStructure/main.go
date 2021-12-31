package main

import (
	"fmt"
	"hash_table/hashtable"
)

func main() {

	myTable := hashtable.New()

	// Test 1
	err := myTable.Insert("test1", 100)
	if err != nil {
		fmt.Println("Error inserting value:", err)
	}

	// Test 2
	exists := myTable.Exists("test1")
	if !exists {
		fmt.Println("1: Value should exist, but doesn't")
	}

	// Test 3
	value, err := myTable.Get("test1")
	if err != nil {
		fmt.Println("2: Unable to get value")
	}
	if value != 100 {
		fmt.Println("3: Value should have been 100, but wasn't")
	}

	// Test 4
	_, err = myTable.Get("test2")
	if err == nil {
		fmt.Println("4: Attempted to get a value that shouldn't exist, but there was no error")
	}

	// Test 5
	err = myTable.Delete("test1")
	if err != nil {
		fmt.Println("5: Error deleting value:", err)
	}

	// Test 6
	_, err = myTable.Get("test1")
	if err == nil {
		fmt.Println("6: Value shouldn't exist anymore, but there was no error")
	}

	// Test 7
	err = myTable.Delete("test2")
	if err == nil {
		fmt.Println("7: Tried to delete a value that shouldn't exist, but there was no error")
	}

}
