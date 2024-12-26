package main // This tells Go that this is the main package, where the program starts

import "fmt" // Import the "fmt" package, which is used for formatting and printing

func printMenu() {
	fmt.Println("1. Increment counter")
	fmt.Println("2. View counter")
	fmt.Println("3. Exit")
	fmt.Println("Choose an option")

}

func main() {
	counter := 0

	// Infinite loop to repeatedly show the menu and handle user choices

	for {
		printMenu()

		var choice int

		_, err := fmt.Scan(&choice) // Scan returns two values(no of inputs and error if any)

		// If there's an error (like invalid input), show an error message
		if err != nil { // nil means 0.
			fmt.Println("Invalid input. Please enter a number.")
			// Flush the input buffer to prevent infinite loop in case of invalid input
			fmt.Scanln()
			continue
		}

		switch choice {
		case 1:
			counter++
			fmt.Println("Counter incremented")
		case 2:
			fmt.Println("Current counter value", counter)

		case 3:
			fmt.Println("Exiting application.")
			return
		default:
			fmt.Println("Invalid option. Please choose again.")
		}

	}
}
