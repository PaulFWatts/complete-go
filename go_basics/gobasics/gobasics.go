package gobasics

import (
	"fmt"

	"github.com/PaulFWatts/complete_go/go_basics/library"
)

type Person struct {
	Name   string
	Age    int
	Active bool
}

// RunBasics demonstrates Go basics concepts.
func RunBasics() {
	fmt.Println("*** Section 1 - Basics ***")
	fmt.Println()
	fmt.Print("Press Enter to continue...")
	fmt.Scanln()

	fmt.Println("Hello, World!")

	// Constants and Variables
	var name string = "John"
	var Myage int = 30
	var isActive bool = true
	fmt.Printf("Name: %s, Age: %d, Active: %v\n", name, Myage, isActive)

	// Short variable declaration
	city := "New York"
	fmt.Printf("City: %s\n", city)

	const pi = 3.14
	fmt.Printf("The value of pi is: %f\n", pi)

	const (
		Sunday = iota + 1
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Printf("The value of Sunday is: %d\n", Sunday)
	fmt.Printf("The value of Monday is: %d\n", Monday)
	fmt.Printf("The value of Tuesday is: %d\n", Tuesday)
	fmt.Printf("The value of Wednesday is: %d\n", Wednesday)
	fmt.Printf("The value of Thursday is: %d\n", Thursday)
	fmt.Printf("The value of Friday is: %d\n", Friday)
	fmt.Printf("The value of Saturday is: %d\n", Saturday)
	fmt.Println()

	// Functions
	result := add(5, 3)
	fmt.Printf("The result of adding 5 and 3 is: %d\n", result)

	sum, product := calulateSumAndProduct(4, 6)
	fmt.Printf("The sum of 4 and 6 is: %d\n", sum)
	fmt.Printf("The product of 4 and 6 is: %d\n", product)
	fmt.Println()

	// Arrays and Slices
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Array: %v\n", arr)
	fmt.Printf("The array size is: %d\n", len(arr))
	fmt.Println("This is the last value", arr[len(arr)-1])
	fmt.Println()

	// Multi-dimensional array
	var matrix [2][3]int = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Printf("Matrix: %v\n", matrix)
	fmt.Println()

	slice := []string{"Go", "is", "awesome"}
	fmt.Printf("Slice: %v\n", slice)

	people := []Person{
		{Name: "Alice", Age: 30, Active: true},
		{Name: "Bob", Age: 22, Active: false},
	}

	for i, person := range people {
		fmt.Printf("Item: %d, %s is %d years old, active: %v\n",
			i, person.Name, person.Age, person.Active)
	}
	fmt.Printf("the slice size is: %d\n", len(people))
	fmt.Println()

	// Slice copy of an array
	original := [5]int{1, 2, 3, 4, 5}
	sliceCopy := original[:]
	fmt.Printf("Original array: %v\n", original)
	fmt.Printf("Slice copy: %v\n", sliceCopy)
	fmt.Printf("The slice capacity is: %d\n", cap(sliceCopy))
	fmt.Println()
	// Append to the slice copy
	sliceCopy = append(sliceCopy, 6)
	fmt.Printf("After appending to slice copy: %v\n", sliceCopy)
	fmt.Printf("Original array after appending to slice copy: %v\n", original)
	fmt.Printf("The slice capacity is: %d\n", cap(sliceCopy))
	fmt.Println()

	// Maps
	personMap := make(map[string]int)
	personMap["Alice"] = 30
	personMap["Bob"] = 22
	fmt.Printf("Person Map: %v\n", personMap)
	ageOfAlice := personMap["Alice"]
	fmt.Printf("Alice's age is: %d\n", ageOfAlice)
	fmt.Println()

	person, exists := personMap["Charlie"]
	if exists {
		fmt.Printf("Charlie's age is: %d\n", person)
	} else {
		fmt.Println("Charlie does not exist in the map.")
	}
	fmt.Println()

	delete(personMap, "Bob")
	fmt.Printf("Person Map after deleting Bob: %v\n", personMap)
	fmt.Println()

	// Structs
	person1 := Person{Name: "Alice", Age: 30, Active: true} // Struct literal
	fmt.Printf("Person1: %v\n", person1)
	fmt.Printf("Person1's name is: %s\n", person1.Name)
	fmt.Printf("Person1's age is: %d\n", person1.Age)
	fmt.Printf("Person1's active status is: %v\n", person1.Active)
	fmt.Println()

	fmt.Print("Press Enter to clear the console...")
	fmt.Scanln()
	library.ClearScreen()
}

func add(a int, b int) int {
	return a + b
}

func calulateSumAndProduct(x int, y int) (int, int) {
	sum := x + y
	product := x * y
	return sum, product
}
