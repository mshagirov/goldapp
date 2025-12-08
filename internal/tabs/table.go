package tabs

import (
	"os"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"

	"github.com/mshagirov/goldap/ldapapi"
)

func NewTable(ti ldapapi.TableInfo) table.Model {
	w, h := GetTableDimensions()
	t := table.New(
		table.WithColumns(ti.Cols),
		table.WithRows(ti.Rows),
		table.WithFocused(true),
		table.WithHeight(h),
		table.WithWidth(w),
		table.WithStyles(GetTableStyle()),
	)
	return t
}

func GetTableStyle() table.Styles {
	s := table.DefaultStyles()
	hlColor := lipgloss.AdaptiveColor{Light: "#0014a8", Dark: "#265ef7"}
	s.Header = s.Header.Foreground(hlColor)
	s.Selected = s.Selected.Foreground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#ffffff"}).Background(hlColor)
	return s
}

func GetTableDimensions() (int, int) {
	termWidth, termHeight, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		termWidth, termHeight = 20, 20
	}
	w, h := windowStyle.GetHorizontalFrameSize(), windowStyle.GetVerticalFrameSize()
	return (termWidth - w), (termHeight - 6*h)
}
