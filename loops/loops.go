package loops

import (
	"fmt"

	"github.com/PaulFWatts/complete_go/library"
)

// RunLoops demonstrates Go loops concepts
func RunLoops() {
	fmt.Println("*** Section 3 - Loops ***")
	fmt.Println()
	fmt.Print("Press Enter to continue...")
	fmt.Scanln()

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
	fmt.Print("Press Enter to clear the console...")
	fmt.Scanln()
	library.ClearScreen()
}
