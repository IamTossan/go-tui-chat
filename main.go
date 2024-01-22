package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gorilla/websocket"
)

const (
	CHAT_URL = "wss://socketsbay.com/wss/v2/1/demo/"
)

func main() {
	c, _, err := websocket.DefaultDialer.Dial(CHAT_URL, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	p := tea.NewProgram(InitialModel(c, CHAT_URL))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
