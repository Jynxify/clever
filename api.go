package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func GetGPTResponse(m Model) string {
	time.Sleep(2 * time.Second)
	return "This is a test response"
}

type apiResponseMsg struct {
	response string
}

// fetchAPI simulates an API call with the given question and context
func fetchAPI(m *Model) tea.Cmd {
	return func() tea.Msg {
		response := GetGPTResponse(*m)
		return apiResponseMsg{response: response}
	}
}
