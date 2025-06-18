package app
import ("github.com/charmbracelet/lipgloss")

var interfaceStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA"))

var servingStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#3079CF"))

var errStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#f55050"))

var logStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#64C964"))
