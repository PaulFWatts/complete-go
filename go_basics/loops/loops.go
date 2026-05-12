package loops

import (
	"fmt"
	"math/rand/v2"

	"github.com/PaulFWatts/complete_go/go_basics/library"
)

// RunLoops demonstrates Go loops concepts.
func RunLoops() {
	fmt.Println("*** Section 3 - Loops ***")
	fmt.Println()
	fmt.Print("Press Enter to continue...")
	if _, err := fmt.Scanln(); err != nil {
		fmt.Println("Unable to read input:", err)
	}

	// Control Structures
	age := rand.IntN(91) + 10
	fmt.Println(classifyAge(age))
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
	if _, err := fmt.Scanln(); err != nil {
		fmt.Println("Unable to read input:", err)
	}
	if err := library.ClearScreen(); err != nil {
		fmt.Println("Unable to clear the console:", err)
	}
}

func classifyAge(age int) string {
	if age < 18 {
		return "You are a minor."
	}
	if age < 65 {
		return "You are an adult."
	}
	return "You are a senior."
}
