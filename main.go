package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	displayServer := detectDisplayServer()
	distribution := getDistributionName()

	if displayServer == "Unknown" {
		fmt.Printf("Error: Could not detect your display server!\n")
		os.Exit(1)
	}

	lux := tea.NewProgram(initialModel(displayServer, distribution))
	if _, err := lux.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
