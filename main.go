package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Active bool
}

func main() {
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

	// Control Structures

	age := 65
	if age < 18 {
		fmt.Println("You are a minor.")
	} else if age >= 18 && age < 65 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a senior.")
	}
	fmt.Println()

	day := "Friday"
	switch day {
	case "Monday":
		fmt.Println("It's the start of the week.")
	case "Friday":
		fmt.Println("It's almost the weekend.")
	default:
		fmt.Println("It's a regular day.")
	}

	for i := 1; i <= 5; i++ {
		fmt.Printf("Iteration %d\n", i)
	}
	fmt.Println()

	// while loop equivalent using for
	count := 1
	for count <= 5 {
		fmt.Printf("Count: %d\n", count)
		count++
	}
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
}

func add(a int, b int) int {
	return a + b
}

func calulateSumAndProduct(x int, y int) (int, int) {
	sum := x + y
	product := x * y
	return sum, product
}
