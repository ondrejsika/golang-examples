package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type step int

const (
	stepName step = iota
	stepPet
	stepPetName
	stepSecret
	stepDone
)

type model struct {
	step     step
	name     string
	pet      string
	petName  string
	secret   string
	cursor   int
	input    string
}

func initialModel() model {
	return model{
		step: stepName,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "enter":
			switch m.step {
			case stepName:
				if m.input != "" {
					m.name = m.input
					m.input = ""
					m.step = stepPet
				}
			case stepPet:
				pets := []string{"Dog", "Cat", "Rat", "Iguana"}
				m.pet = pets[m.cursor]
				m.step = stepPetName
			case stepPetName:
				if m.input != "" {
					m.petName = m.input
					m.input = ""
					m.step = stepSecret
				}
			case stepSecret:
				m.secret = m.input
				m.step = stepDone
				return m, tea.Quit
			}

		case "up", "k":
			if m.step == stepPet && m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.step == stepPet && m.cursor < 3 {
				m.cursor++
			}

		case "backspace":
			if m.step == stepName || m.step == stepPetName || m.step == stepSecret {
				if len(m.input) > 0 {
					m.input = m.input[:len(m.input)-1]
				}
			}

		default:
			if m.step == stepName || m.step == stepPetName || m.step == stepSecret {
				m.input += msg.String()
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	var s string

	switch m.step {
	case stepName:
		s = "What is your name?\n\n"
		s += fmt.Sprintf("> %s\n\n", m.input)
		s += "Press enter to continue"

	case stepPet:
		s = "Choose your favorite pet:\n\n"
		choices := []string{"Dog", "Cat", "Rat", "Iguana"}
		for i, choice := range choices {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}
		s += "\nUse arrow keys to select, enter to confirm"

	case stepPetName:
		s = fmt.Sprintf("What is your %s's name?\n\n", m.pet)
		s += fmt.Sprintf("> %s\n\n", m.input)
		s += "Press enter to continue"

	case stepSecret:
		s = "Enter your secret:\n\n"
		masked := ""
		for range m.input {
			masked += "*"
		}
		s += fmt.Sprintf("> %s\n\n", masked)
		s += "Press enter to continue"

	case stepDone:
		s = "Thank you!\n"
	}

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	m, err := p.Run()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	if finalModel, ok := m.(model); ok && finalModel.step == stepDone {
		fmt.Printf("\nName: %s\n", finalModel.name)
		fmt.Printf("Favorite Pet: %s\n", finalModel.pet)
		fmt.Printf("Pet Name: %s\n", finalModel.petName)
		fmt.Printf("Secret: %s\n", finalModel.secret)
	}
}
