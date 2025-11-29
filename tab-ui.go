package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mshagirov/goldap/internal/tabs"
)

func runTabs(tabNames []string, tabContent []string) {

	m := tabs.Model{TabNames: tabNames, TabContent: tabContent}
	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
