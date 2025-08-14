package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func detectDisplayServer() string {
	if os.Getenv("WAYLAND_DISPLAY") != "" {
		return "Wayland"
	}
	if os.Getenv("DISPLAY") != "" {
		return "X11"
	}
	return "Unknown"
}

func getDistributionName() string {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return "Unknown"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "NAME=") {
			return strings.Trim(strings.Split(line, "=")[1], "\"")
		}
	}

	return "Unknown"
}

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
