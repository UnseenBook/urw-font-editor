package editor

import (
	"github.com/charmbracelet/lipgloss"
	"strings"
)

var PixelStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#222222")).
	Foreground(lipgloss.Color("#dddddd"))

var Selected = lipgloss.NewStyle().
	Background(lipgloss.Color("#333333")).
	Foreground(lipgloss.Color("#ffffff"))

var CharStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.InnerHalfBlockBorder()).
	BorderForeground(lipgloss.Color("1"))

func (p Pixel) String() string {
	text := "  "
	if p {
		text = "██"
	}
	return PixelStyle.Render(text)
}

func (c Char) String() string {
	line := make([]string, 0, c.Width())
	lines := make([]string, 0, c.Height())
	for _, pixels := range c {
		for _, pixel := range pixels {
			line = append(line, pixel.String())
		}
		lines = append(lines, strings.Join(line, ""))
		line = line[:0]
	}
	return CharStyle.Render(lipgloss.JoinVertical(0, lines...))
}

func (c Char) Width() int {
	return len(c[0])
}

func (c Char) Height() int {
	return len(c)
}

func (f Font) String() string {
	line := make([]string, 0, 1)
	lines := make([]string, 0, 1)
	for i, char := range f.Chars {
		if i != 0 && i%f.Width == 0 {
			lines = append(lines, lipgloss.JoinHorizontal(lipgloss.Top, line...))
			line = line[:0]
		}
		line = append(line, char.String())
	}
	lines = append(lines, lipgloss.JoinHorizontal(lipgloss.Top, line...))

	return lipgloss.JoinVertical(0, lines...)
}
