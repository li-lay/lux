package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	logo     string
	desc     string
	protocol string
	distro   string
}

func initialModel(displaySever string, distribution string) model {
	return model{
		logo: `
.____
|    |    __ _____  ___
|    |   |  |  \  \/  /
|    |___|  |  />    < 
|_______ \____//__/\_ \
        \/           \/
`,
		desc:     "Very Simple Brightness Controller\n\nMeaning: Lux - Latin word for Light\n",
		protocol: displaySever,
		distro:   distribution,
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

	content += Yellow.Render("\n\nDetected Monitors:\n")

	content += Grey.Render("\nPress q to quit.\n") + "\n"
	return content
}
