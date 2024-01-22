package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/gorilla/websocket"
)

type Model struct {
	textInput textinput.Model
	error     error
	sub       chan struct{}
	responses []string
	conn      *websocket.Conn
	message   chan string
	url       string
}
type ErrMsg struct{ error }

type ChatMsg string

func InitialModel(conn *websocket.Conn, url string) Model {
	ti := textinput.New()
	ti.Placeholder = "Write something and press enter to send a message"
	ti.Focus()
	ti.CharLimit = 70
	ti.Width = 78

	return Model{
		textInput: ti,
		sub:       make(chan struct{}),
		conn:      conn,
		message:   make(chan string),
		url:       url,
	}
}

type ResponseMsg string
