// Package lipglossstyles
package lipglossstyles

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var DirStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#7D56F4")).
	MarginRight(1)

var UsernameStyle = lipgloss.NewStyle().
	Bold(false).
	Foreground(lipgloss.Color("#9c7ef7"))

func RenderHeader(dir, username string) {
	usernameString := fmt.Sprintf("%s > ", username)

	usernameRender := UsernameStyle.Render(usernameString)
	dirRender := DirStyle.Render(dir)

	fmt.Println(dirRender + usernameRender)
}
