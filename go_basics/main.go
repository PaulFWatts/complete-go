package main

import (
	"fmt"

	"github.com/PaulFWatts/complete_go/go_basics/gobasics"
	"github.com/PaulFWatts/complete_go/go_basics/loops"
	"github.com/PaulFWatts/complete_go/go_basics/structs"
)

// main runs the Go_Basics app with a login menu.
func main() {
	for {
		displayMenu()
		choice := getUserChoice()

		switch choice {
		case 1:
			gobasics.RunBasics()
		case 2:
			loops.RunLoops()
		case 3:
			structs.RunStructs()
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

// displayMenu shows the login/main menu.
func displayMenu() {
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("       Welcome to Go Basics Demo")
	fmt.Println("========================================")
	fmt.Println("1. Run Basics")
	fmt.Println("2. Run Loops")
	fmt.Println("3. Run Structs")
	fmt.Println("4. Exit")
	fmt.Println("========================================")
}

// getUserChoice reads and returns the user's menu selection.
func getUserChoice() int {
	var choice int
	fmt.Print("Enter your choice (1-4): ")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return 0
	}
	return choice
}
