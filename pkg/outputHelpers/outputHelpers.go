package outputHelpers

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func PrettyPrint(s string) {
	var anotherStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		Background(lipgloss.Color("#ff79c6")).
		Foreground(lipgloss.Color("#000000")).
		BorderBackground(lipgloss.Color("#ff79c6")).
		BorderForeground(lipgloss.Color("#000000")).
		BorderRight(true).
		BorderLeft(true).
		BorderTop(true).
		BorderBottom(true)

	fmt.Println(anotherStyle.Render(s))
}
