package tui

// code from bubletea's example

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor int
	choice chan string
	list   []string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			close(m.choice) // If we're quitting just close the channel.
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.choice <- m.list[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(m.list) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.list) - 1
			}
		}

	}

	return m, nil
}

func (m model) View() string {
	s := strings.Builder{}
	s.WriteString("Select server to connect\n\n")

	for i := 0; i < len(m.list); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		if len(strings.Split(m.list[i], "\t"))>1{
			s.WriteString(strings.Split(m.list[i], "\t")[1])
		}else {
			s.WriteString(m.list[i])
		}
		s.WriteString("\n")
	}

	return s.String()
}

func RunTUI(list []string) (string, error) {
	// This is where we'll listen for the choice the user makes in the Bubble
	// Tea program.
	result := make(chan string, 1)

	// Pass the channel to the initialize function so our Bubble Tea program
	// can send the final choice along when the time comes.
	p := tea.NewProgram(model{cursor: 0, choice: result, list: list})
	if err := p.Start(); err != nil {
		return "", err
	}

	// Print out the final choice.
	if r := <-result; r != "" {
		return strings.Split(r,"\t")[0], nil
	}

	return "", nil
}
