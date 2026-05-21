package main

import "fmt"

// Item struct
type Item struct {
	Name string
	Type string
}

// Player struct
type Player struct {
	Name      string
	Inventory []Item
}

func main() {
	fmt.Println("========================================")
	fmt.Println("       Welcome to Role Play Exercise")
	fmt.Println("========================================")

	// Get player name
	fmt.Print("Enter your character name: ")
	var playerName string
	_, err := fmt.Scanln(&playerName)
	if err != nil {
		playerName = "Adventurer"
	}

	// Create player
	player := &Player{
		Name:      playerName,
		Inventory: []Item{},
	}

	fmt.Printf("\nWelcome, %s!\n\n", player.Name)

	// Game loop
	for {
		displayGameMenu()
		choice := getGameChoice()

		switch choice {
		case 1:
			pickUpItemMenu(player)
		case 2:
			dropItemMenu(player)
		case 3:
			useItemMenu(player)
		case 4:
			viewInventory(player)
		case 5:
			fmt.Printf("\nGoodbye, %s!\n", player.Name)
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

// displayGameMenu shows the game menu options.
func displayGameMenu() {
	fmt.Println("\n========================================")
	fmt.Println("                  Menu")
	fmt.Println("========================================")
	fmt.Println("1. Pick Up Item")
	fmt.Println("2. Drop Item")
	fmt.Println("3. Use Item")
	fmt.Println("4. View Inventory")
	fmt.Println("5. Exit")
	fmt.Println("========================================")
}

// getGameChoice reads the user's menu selection.
func getGameChoice() int {
	var choice int
	fmt.Print("Enter your choice (1-5): ")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		return 0
	}
	return choice
}

// pickUpItemMenu handles picking up an item.
func pickUpItemMenu(p *Player) {
	fmt.Print("Enter item name: ")
	var itemName string
	_, err := fmt.Scanln(&itemName)
	if err != nil {
		return
	}

	fmt.Print("Enter item type (weapon/armor/potion): ")
	var itemType string
	_, err = fmt.Scanln(&itemType)
	if err != nil {
		return
	}

	item := Item{Name: itemName, Type: itemType}
	p.PickUpItem(item)
}

// dropItemMenu handles dropping an item.
func dropItemMenu(p *Player) {
	if len(p.Inventory) == 0 {
		fmt.Println("Your inventory is empty!")
		return
	}

	fmt.Print("Enter item name to drop: ")
	var itemName string
	_, err := fmt.Scanln(&itemName)
	if err != nil {
		return
	}

	p.DropItem(itemName)
}

// useItemMenu handles using an item.
func useItemMenu(p *Player) {
	if len(p.Inventory) == 0 {
		fmt.Println("Your inventory is empty!")
		return
	}

	fmt.Print("Enter item name to use: ")
	var itemName string
	_, err := fmt.Scanln(&itemName)
	if err != nil {
		return
	}

	p.UseItem(itemName)
}

// viewInventory displays the player's current inventory.
func viewInventory(p *Player) {
	fmt.Printf("\n%s's Inventory:\n", p.Name)
	if len(p.Inventory) == 0 {
		fmt.Println("  (empty)")
		return
	}
	for i, item := range p.Inventory {
		fmt.Printf("  %d. %s (%s)\n", i+1, item.Name, item.Type)
	}
}

// PickUpItem Pick up an item (modifies the player's inventory)
func (p *Player) PickUpItem(item Item) {
	p.Inventory = append(p.Inventory, item)
	fmt.Printf("%s picked up %s!\n", p.Name, item.Name)

}

// DropItem Drop an item (removes from inventory)
func (p *Player) DropItem(itemName string) {
	for i, item := range p.Inventory {
		if item.Name == itemName {
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			fmt.Printf("%s dropped %s!\n", p.Name, itemName)
			return
		}
	}
	fmt.Printf("%s does not have %s in inventory.\n", p.Name, itemName)
}

// UseItem Use an item (if potion, remove it after use)
func (p *Player) UseItem(itemName string) {
	for i, item := range p.Inventory {
		if item.Name == itemName {
			fmt.Printf("%s used %s!\n", p.Name, item.Name)
			if item.Type == "potion" {
				p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
				fmt.Printf("%s consumed the potion and it is now removed from inventory.\n", p.Name)
			}
			return
		}
	}
	fmt.Printf("%s does not have %s in inventory.\n", p.Name, itemName)
}
