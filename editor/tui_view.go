package editor

import (
	"github.com/charmbracelet/lipgloss"
)

var PixelStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#333333")).
	Foreground(lipgloss.Color("#ffffff"))

var CharStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.InnerHalfBlockBorder()).
	BorderForeground(lipgloss.Color("#a0a0a0"))

func (p Pixel) String() string {
	text := "  "
	if p {
		text = "██"
	}
	return PixelStyle.Render(text)
}

func (c Char) String() string {
	text := ""
	lines := make([]string, 0, 10)
	for _, pixels := range c {
		for _, pixel := range pixels {
			text += pixel.String()
		}
		lines = append(lines, text)
		text = ""
	}
	return CharStyle.Render(lipgloss.JoinVertical(0, lines...))
}

func (f Font) String() string {
	text := ""
	lines := make([]string, 0, 16)
	for i, char := range f.Chars {
		if i != 0 && i%16 == 0 {
			lines = append(lines, text)
			text = ""
		}
		if text == "" {
			text = char.String()
			continue
		}
		text = lipgloss.JoinHorizontal(0, text, char.String())
	}
	lines = append(lines, text)

	return lipgloss.JoinVertical(0, lines...)
}
