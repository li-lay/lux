# Lux - Very Simple Brightness Controller

Lux is a Terminal User Interface (TUI) application use to controll monitor brightness on Linux distributions. The name "Lux" comes from the Latin word for Light.

## Features

- Detects your Linux distribution
- Detects your display server (X11 or Wayland)
- Clean, colorful terminal interface thanks to Bubbletea and Lipgloss

## Installation

To install and run Lux, you need to have Go installed on your system:

```bash
# Clone the repository
git clone <repository-url>
cd lux

# Build the application
go build

# Run the application
./lux
```

Alternatively, you can run it directly without building:

```bash
go run .
```

## Usage

Once launched, Lux will display system information in a terminal interface. Use the following keys to interact with the application:

- `q` or `Ctrl+C`: Quit the application

The interface shows:

1. A stylized "Lux" logo
2. Application description
3. Operating system information
4. Display server protocol
5. Detected monitors (planned feature)

## Dependencies

Lux uses the following Go libraries:

- [Bubbletea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Styling and layout

## Development

### Build Commands

- Build: `go build`
- Run: `go run .`

### Code Style

- Follow Go's standard formatting (gofmt)
- Use descriptive names for variables and functions
- Handle errors explicitly
- Group standard library imports separately from third-party imports

## Contributing

[Contribution guidelines would go here]

This project is a work in progress and aims to become a simple brightness controller for Linux systems.

