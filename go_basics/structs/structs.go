package structs

import (
	"fmt"

	"github.com/PaulFWatts/complete_go/go_basics/library"
)

type Person struct {
	Name string
	Age  int
}

// Address represents a postal address used in nested structs
type Address struct {
	Street string
	City   string
	Zip    string
}

// Employee demonstrates nesting a named struct (Person) and another struct (Address)
type Employee struct {
	Info    Person
	ID      int
	Home    Address
	Manager *Person
}

// RunStructs demonstrates Go struct concepts.
func RunStructs() {
	fmt.Println("*** Section 4 - Structs ***")
	fmt.Println()
	fmt.Print("Press Enter to continue...")
	if _, err := fmt.Scanln(); err != nil {
		fmt.Println("Unable to read input:", err)
	}

	// Named Struct
	person := Person{Name: "Paul", Age: 69}
	fmt.Printf("Person struct: %+v\n", person)
	fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
	fmt.Println()

	// Anonymous Struct
	employee := struct {
		name string
		id   int
	}{
		name: "Bob", id: 1}
	fmt.Printf("Employee struct: %+v\n", employee)
	fmt.Printf("Name: %s, ID: %d\n", employee.name, employee.id)
	fmt.Println()

	// Nested Structs
	// Create an Address
	addr := Address{Street: "1 Infinite Loop", City: "Cupertino", Zip: "95014"}

	// Create an Employee that nests Person and Address
	emp := Employee{
		Info:    Person{Name: "Alice", Age: 34},
		ID:      42,
		Home:    addr,
		Manager: &Person{Name: "Bob", Age: 45},
	}

	fmt.Printf("Employee (nested structs): %+v\n", emp)
	fmt.Printf("Employee Name: %s, Age: %d, ID: %d\n", emp.Info.Name, emp.Info.Age, emp.ID)
	fmt.Printf("Employee Home: %s, %s %s\n", emp.Home.Street, emp.Home.City, emp.Home.Zip)
	fmt.Printf("Manager: %s (age %d)\n", emp.Manager.Name, emp.Manager.Age)
	fmt.Println()

	fmt.Print("Press Enter to clear the console...")
	if _, err := fmt.Scanln(); err != nil {
		fmt.Println("Unable to read input:", err)
	}
	if err := library.ClearScreen(); err != nil {
		fmt.Println("Unable to clear the console:", err)
	}
}
