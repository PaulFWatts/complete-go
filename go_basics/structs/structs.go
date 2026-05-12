package structs

import (
	"fmt"

	"github.com/PaulFWatts/complete_go/go_basics/library"
)

// RunStructs demonstrates Go struct concepts.
func RunStructs() {
	fmt.Println("*** Section 4 - Structs ***")
	fmt.Println()
	fmt.Print("Press Enter to continue...")
	if _, err := fmt.Scanln(); err != nil {
		fmt.Println("Unable to read input:", err)
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
