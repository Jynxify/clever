package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// Model is the bubbletea model
type Model struct {
	// Input is the text input model
	UserInput        textinput.Model
	Context          map[string]string
	PreviousMessages []map[string]string
	Question         string
}

func initialModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Enter your query here"
	ti.Focus()
	ti.CharLimit = 300

	return Model{
		UserInput:        ti,
		Context:          make(map[string]string),
		PreviousMessages: []map[string]string{},
		Question:         "",
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

// Update updates the bubbletea model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.UserInput.SetValue(msg.String())
		if msg.String() == "ctrl+enter" {
			if strings.TrimSpace(m.UserInput.Value()) != "" {
				// m.UserInput.SetValue("Command executed !")
				m.UserInput.CursorEnd()
				m.UserInput.Focus()
				return m, nil
			} else {
				// m.UserInput.SetValue("Something went wrong !")
				m.UserInput.CursorEnd()
				m.UserInput.Focus()
				return m, nil
			}
		}

		// Handle other key presses
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			m.UserInput.SetValue("")
		case "enter":
			// Get the user input
			if strings.TrimSpace(m.UserInput.Value()) != "" {
				// userInput := m.UserInput.Value()
				// m.Question = userInput
				m.UserInput.Blur()
				return m, fetchAPI(&m)
			}
		}

	case apiResponseMsg:
		// Update user input with API response
		m.UserInput.SetValue(msg.response)
		m.UserInput.CursorEnd()
		m.UserInput.Focus()
		return m, nil
	}

	// Pass the message to the input component
	userInput, cmd := m.UserInput.Update(msg)
	m.UserInput = userInput

	return m, cmd
}

// View returns the bubbletea view
func (m Model) View() string {
	return m.UserInput.View()
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error starting program:", err)
	}
}
