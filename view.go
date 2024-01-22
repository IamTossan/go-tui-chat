package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	textInputStyle = lipgloss.NewStyle().Width(80).Padding(1, 2).Margin(1, 0).BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("63"))
	messagesStyle  = lipgloss.NewStyle().Width(160).Padding(1, 2).Margin(1, 0).BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("63"))
)

func (m Model) View() string {
	s := ""
	if m.error != nil {
		s += fmt.Sprintf("\nSomething went wrong: %s\n", m.error)
	}

	messages := ""
	messages += fmt.Sprintf("Connected to: %s\n", m.url)
	for _, message := range m.responses {
		messages += fmt.Sprintf("\n> %s\n", message)
	}
	s += messagesStyle.Render(messages)

	s += textInputStyle.Render(m.textInput.View() + "\n\nPress esc or ctrl+c to exit")

	return s
}
