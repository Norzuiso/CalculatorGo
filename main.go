package main

import (
	"fmt"
)

func main() {
	showMenu()
}

func showMenu() {
	var option int
	for option != 6 {
		menu := `
	Welcome!
	[1] addition.
	[2] subtraction.
	[3] multiplication.
	[4] division.
	[5] Rules and instructions.
	[6] exit.`
		fmt.Println(menu)
		fmt.Scan(&option)
		options(option)
	}
}

func options(option int) {
	switch option {
	case 1:
		fmt.Printf("Additon.")
		break
	case 2:
		fmt.Printf("Substrction.")
		break
	case 3:
		fmt.Printf("Multiplication.")
		break
	case 4:
		fmt.Printf("Division.")
		break
	case 5:
		fmt.Printf("Rules and instructions.")
		break
	case 6:
		fmt.Printf("Exit")
		break
	}
}
