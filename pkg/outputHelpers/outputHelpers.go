package outputHelpers

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)


func PrettyPrint(s string) {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#44475a")).
		Background(lipgloss.Color("#ff79c6")).
		PaddingTop(0).
		PaddingLeft(4).
		Width(80)

	fmt.Println(style.Render("Hello, kitty"))
}
