package main

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
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

type Monitor struct {
	Output  string
	Make    string
	Model   string
	Serial  string
	Enabled bool
}

func getMonitorsFromWayland() ([]Monitor, error) {
	cmd := exec.Command("wlr-randr")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	var monitors []Monitor

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if !strings.HasPrefix(line, "  ") && strings.Contains(line, " ") && len(strings.TrimSpace(line)) > 0 {
			parts := strings.SplitN(strings.TrimSpace(line), " ", 2)
			if len(parts) >= 1 {
				monitor := Monitor{Output: parts[0]}
				for j := i + 1; j < len(lines); j++ {
					if lines[j] == "" {
						break
					}
					if strings.HasPrefix(lines[j], "  Make: ") {
						monitor.Make = strings.TrimSpace(strings.TrimPrefix(lines[j], "  Make: "))
					} else if strings.HasPrefix(lines[j], "  Model: ") {
						monitor.Model = strings.TrimSpace(strings.TrimPrefix(lines[j], "  Model: "))
					} else if strings.HasPrefix(lines[j], "  Serial: ") {
						monitor.Serial = strings.TrimSpace(strings.TrimPrefix(lines[j], "  Serial: "))
					} else if strings.HasPrefix(lines[j], "  Enabled: ") {
						enabled := strings.TrimSpace(strings.TrimPrefix(lines[j], "  Enabled: "))
						monitor.Enabled = (enabled == "yes")
					}
				}
				monitors = append(monitors, monitor)
			}
		}
	}

	return monitors, nil
}
