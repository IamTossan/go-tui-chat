package main

import (
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gorilla/websocket"
)

func subscribeToMessages(sub chan string, conn *websocket.Conn) tea.Cmd {
	return func() tea.Msg {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return ErrMsg{error: err}
			}
			sub <- string(message)
		}
	}
}

func (m Model) sendMessage() {
	w, err := m.conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return
	}
	w.Write([]byte(m.textInput.Value()))
	w.Close()
}

// A command that waits for the activity on a channel.
func waitForMessage(sub chan string) tea.Cmd {
	return func() tea.Msg {
		return ResponseMsg(<-sub)
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		textinput.Blink,
		subscribeToMessages(m.message, m.conn),
		waitForMessage(m.message),
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			m.responses = append(m.responses, m.textInput.Value())
			m.sendMessage()
			m.textInput.Reset()
			return m, nil
		}
	case ErrMsg:
		m.error = msg
		return m, nil
	case ResponseMsg:
		m.responses = append(m.responses, string(msg))
		return m, waitForMessage(m.message) // wait for next event
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}
