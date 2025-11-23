package login

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).PaddingTop(2).PaddingLeft(4)
	blurredStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).PaddingLeft(4)

	cursorStyle        = focusedStyle
	noStyle            = lipgloss.NewStyle().PaddingTop(2).PaddingLeft(4)
	titleText   string = `LDAP Administrator Password`
	titleSyle          = blurredStyle.PaddingTop(2).Bold(true)
	helpStyle          = blurredStyle
	helpText    string = `(press esc or ctrl+c to exit)`

	focusedButton = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).PaddingLeft(4).
			Render("[ Submit ]")
	blurredButton = fmt.Sprintf("%s",
		lipgloss.NewStyle().Foreground(
			lipgloss.Color("240")).PaddingLeft(4).
			Render("[ Submit ]"))
)

type model struct {
	focusIndex int
	inputs     []textinput.Model
}

func initialModel() model {
	m := model{
		inputs: make([]textinput.Model, 1),
	}

	var t textinput.Model
	t = textinput.New()
	t.Cursor.Style = cursorStyle
	t.CharLimit = 32

	// placeholder problem: displays only 1st char
	t.Placeholder = "_____"
	t.EchoMode = textinput.EchoPassword
	t.EchoCharacter = '•'
	t.PromptStyle = focusedStyle
	t.TextStyle = focusedStyle
	t.Focus()

	m.inputs[0] = t

	return m
}

func (m model) Init() tea.Cmd {
	return tea.Batch(tea.SetWindowTitle("goldap | login"), textinput.Blink)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			if s == "enter" && m.focusIndex == len(m.inputs) {
				return m, tea.Quit
			}

			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m model) View() string {
	var b strings.Builder

	b.WriteString(titleSyle.Render(titleText))

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	b.WriteString(helpStyle.Render(helpText))
	return b.String()
}

func Run() (string, error) {
	m := initialModel()
	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		return "", fmt.Errorf("could not start program: %s", err)
	}
	return m.inputs[0].Value(), nil
}
