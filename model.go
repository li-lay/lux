package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	logo     string
	desc     string
	protocol string
	distro   string
	monitors []Monitor
}

func initialModel(displaySever string, distribution string) model {
	monitors, err := getMonitorsFromWayland()
	if err != nil {
		monitors = []Monitor{}
	}

	return model{
		logo: `
.____
|    |    __ _____  ___
|    |   |  |  \  \/  /
|    |___|  |  />    < 
|_______ \____//__/\_ \
         \/           \/
`,
		desc:     "Very Simple Brightness Controller\n",
		protocol: displaySever,
		distro:   distribution,
		monitors: monitors,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	content := Green.Render(m.logo) + "\n"
	content += Fg.Render(m.desc) + "\n"

	content += Orange.Render("Informations" + "\nOS: " + m.distro + "\nProtocol: " + m.protocol)

	content += Yellow.Render("\n\nDetected Monitors:")

	// Display monitors
	if len(m.monitors) == 0 {
		content += Grey.Render("No monitors detected\n")
	} else {
		for i, monitor := range m.monitors {
			monitorInfo := fmt.Sprintf("\n%d. %s %s", i+1, Blue.Render(monitor.Make), Blue.Render(monitor.Model))
			if !monitor.Enabled {
				monitorInfo += " (disabled)"
			}
			content += Fg.Render(monitorInfo + "\n")
		}
	}

	content += Grey.Render("\nPress q to quit.\n") + "\n"
	return content
}
