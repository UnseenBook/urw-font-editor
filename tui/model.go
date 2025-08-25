package tui

import (
	"fmt"
	"strings"

	"github.com/UnseenBook/urw-font-editor/editor"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

var (
	SelectedPixelBorderColor  = lipgloss.Color("4")
	HighlightPixelColor       = lipgloss.Color("5")
	PixelOnColor              = lipgloss.Color("15")
	PixelOnWarningColor       = lipgloss.Color("1")
	PixelOffColor             = lipgloss.Color("0")
	SelectedCharBorderColor   = lipgloss.Color("7")
	UnSelectedCharBorderColor = lipgloss.Color("8")
	SelectedPixelBorder       = lipgloss.ThickBorder()
	PixelOnStyle              = lipgloss.NewStyle().
					Foreground(SelectedPixelBorderColor).
					Background(PixelOnColor)
	PixelOnWarningStyle = lipgloss.NewStyle().
				Foreground(SelectedPixelBorderColor).
				Background(PixelOnWarningColor)
	PixelOffStyle = lipgloss.NewStyle().
			Foreground(SelectedPixelBorderColor).
			Background(PixelOffColor)
	UnselectedCharStyle = lipgloss.NewStyle().
				BorderStyle(lipgloss.InnerHalfBlockBorder()).
				BorderForeground(UnSelectedCharBorderColor)
	SelectedCharStyle = lipgloss.NewStyle().
				BorderStyle(lipgloss.InnerHalfBlockBorder()).
				BorderForeground(SelectedCharBorderColor)
)

func NewFontViewer(f editor.Font) FontViewer {
	return FontViewer{font: f, selectedChar: 2, selectedPixel: Vector{2, 2}}
}

type Vector struct {
	x, y int
}

type FontViewer struct {
	ready         bool
	font          editor.Font
	viewport      viewport.Model
	selectedChar  int
	selectedPixel Vector
	dimensions    tea.WindowSizeMsg
}

func (m FontViewer) Name() string {
	return "Font Viewer"
}

func (m FontViewer) Init() tea.Cmd {
	return nil
}

func (m FontViewer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.MouseMsg:
		if msg.Action != tea.MouseActionRelease || msg.Button != tea.MouseButtonLeft {
			return m, nil
		}
		if zone.Get("selectedChar").InBounds(msg) {
			posX, posY := zone.Get("selectedChar").Pos(msg)
			m.selectedPixel = Vector{posX / 2, posY}
			m.font.Chars[m.selectedChar][m.selectedPixel.y][m.selectedPixel.x] = !m.font.Chars[m.selectedChar][m.selectedPixel.y][m.selectedPixel.x]
			return m, nil
		}
		top, right, bottom, left := viewableCharsRange(viewUnselectedChar(m.font.Chars[m.selectedChar], m.selectedPixel), m.dimensions)
		charX := m.selectedChar % m.font.Width
		charY := m.selectedChar / m.font.Width
		for y := max(charY-top, 0); y <= charY+bottom && y < m.font.Height; y++ {
			for x := max(charX-left, 0); x <= charX+right && x < m.font.Width; x++ {
				if zone.Get(fmt.Sprintf("char:x%d,y%d", x, y)).InBounds(msg) {
					m.selectedChar = y*m.font.Width + x
					return m, nil
				}
			}
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "ctrl+right":
			// If selectedChar is at the edge don't do anything
			if m.selectedChar%m.font.Width == m.font.Width-1 {
				return m, nil
			}
			m.selectedChar++
			return m, nil
		case "ctrl+left":
			// If selectedChar is at the left edge don't do anything
			if m.selectedChar%m.font.Width == 0 {
				return m, nil
			}
			// Move to the left Char
			m.selectedChar--
			return m, nil
		case "ctrl+down":
			// If at the bottom of the font
			if m.selectedChar/m.font.Width == m.font.Height-1 {
				return m, nil
			}
			m.selectedChar += m.font.Width
			return m, nil
		case "ctrl+up":
			if m.selectedChar/m.font.Width == 0 {
				return m, nil
			}
			m.selectedChar -= m.font.Width
			return m, nil
		case "right":
			// If cursor is at the char right edge
			if m.selectedPixel.x == m.font.Chars[m.selectedChar].Width()-1 {
				// If selectedChar is at the edge don't do anything
				if m.selectedChar%m.font.Width == m.font.Width-1 {
					return m, nil
				}
				// Move to the right Char
				m.selectedPixel.x = 0
				m.selectedChar++
				return m, nil
			}
			// Can move right within the current Char
			m.selectedPixel.x++
			return m, nil
		case "left":
			// If cursor is at the char left edge
			if m.selectedPixel.x == 0 {
				// If selectedChar is at the left edge don't do anything
				if m.selectedChar%m.font.Width == 0 {
					return m, nil
				}
				// Move to the left Char
				m.selectedChar--
				m.selectedPixel.x = m.font.Chars[m.selectedChar].Width() - 1
				return m, nil
			}
			// Can move left withing the current Char
			m.selectedPixel.x--
			return m, nil
		case "down":
			// If cursor is at the char bottom
			if m.selectedPixel.y == m.font.Chars[m.selectedChar].Height()-1 {
				// If at the bottom of the font
				if m.selectedChar/m.font.Width == m.font.Height-1 {
					return m, nil
				}
				m.selectedPixel.y = 0
				m.selectedChar += m.font.Width
				return m, nil
			}
			m.selectedPixel.y++
		case "up":
			if m.selectedPixel.y == 0 {
				if m.selectedChar/m.font.Width == 0 {
					return m, nil
				}
				m.selectedChar -= m.font.Width
				m.selectedPixel.y = m.font.Chars[m.selectedChar].Height() - 1
				return m, nil
			}
			m.selectedPixel.y--
		case " ":
			m.font.Chars[m.selectedChar][m.selectedPixel.y][m.selectedPixel.x] = !m.font.Chars[m.selectedChar][m.selectedPixel.y][m.selectedPixel.x]
		}
		return m, nil
	case tea.WindowSizeMsg:
		m.dimensions = msg
		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height)
			m.viewport.SetContent(m.font.String())
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height
		}
	}
	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func viewUnselectedChar(c editor.Char, highlight Vector) string {
	s := strings.Builder{}
	for y, pixelLine := range c {
		for x, pixel := range pixelLine {
			style := PixelOffStyle
			if pixel {
				style = PixelOnStyle
				if y > 7 {
					style = PixelOnWarningStyle
				}
			}
			ps := "  "
			if x == highlight.x && y == highlight.y {
				style = style.Foreground(HighlightPixelColor)
				ps = "<>"
			}
			s.WriteString(style.Render(ps))
		}
		if y != len(c)-1 {
			s.WriteRune('\n')
		}
	}
	return UnselectedCharStyle.Render(s.String())
}

func viewSelectedChar(c editor.Char, selectedPixel Vector) string {
	segments := make([]string, 0, 3)
	top := strings.Builder{}
	left := strings.Builder{}
	right := strings.Builder{}
	bottom := strings.Builder{}
	topPixels := c[0:selectedPixel.y]
	leftPixelLine := c[selectedPixel.y][:selectedPixel.x]
	rightPixelLine := c[selectedPixel.y][selectedPixel.x+1:]
	border := SelectedPixelBorder

	bottomPixels := c[selectedPixel.y+1:]

	for y, pixelLine := range topPixels {
		for x, pixel := range pixelLine {
			style := PixelOffStyle
			if pixel {
				style = PixelOnStyle
				if y > 7 {
					style = PixelOnWarningStyle
				}
			}
			pixStr := "  "
			switch {
			case y == selectedPixel.y-1 && x == selectedPixel.x-1:
				pixStr = fmt.Sprintf(" %s", border.TopLeft)
			case y == selectedPixel.y-1 && x == selectedPixel.x:
				pixStr = fmt.Sprintf("%s%s", border.Top, border.Top)
			case y == selectedPixel.y-1 && x == selectedPixel.x+1:
				pixStr = fmt.Sprintf("%s ", border.TopRight)
			}
			top.WriteString(style.Render(pixStr))
		}
		if y != len(topPixels)-1 {
			top.WriteRune('\n')
		}
	}
	if top.Len() > 0 {
		segments = append(segments, top.String())
	}

	for x, pixel := range leftPixelLine {
		style := PixelOffStyle
		if pixel {
			style = PixelOnStyle
			if selectedPixel.y > 7 {
				style = PixelOnWarningStyle
			}
		}
		pixStr := "  "
		if x == len(leftPixelLine)-1 {
			pixStr = fmt.Sprintf(" %s", border.Left)
		}

		left.WriteString(style.Render(pixStr))
	}
	selected := c[selectedPixel.y][selectedPixel.x]

	selectedStyle := PixelOffStyle
	if selected {
		selectedStyle = PixelOnStyle
		if selectedPixel.y > 7 {
			selectedStyle = PixelOnWarningStyle
		}
	}
	dot := selectedStyle.Render("  ")

	for x, pixel := range rightPixelLine {
		style := PixelOffStyle
		if pixel {
			style = PixelOnStyle
			if selectedPixel.y > 7 {
				style = PixelOnWarningStyle
			}
		}
		pixStr := "  "
		if x == 0 {
			pixStr = fmt.Sprintf("%s ", border.Right)
		}
		right.WriteString(style.Render(pixStr))
	}
	segments = append(segments, lipgloss.JoinHorizontal(lipgloss.Top, left.String(), dot, right.String()))
	for yp, pixelLine := range bottomPixels {
		for x, pixel := range pixelLine {
			style := PixelOffStyle
			if pixel {
				style = PixelOnStyle
				if yp+1+selectedPixel.y > 7 {
					style = PixelOnWarningStyle
				}
			}
			pixStr := "  "
			switch {
			case yp == 0 && x == selectedPixel.x-1:
				pixStr = fmt.Sprintf(" %s", border.BottomLeft)
			case yp == 0 && x == selectedPixel.x:
				pixStr = fmt.Sprintf("%s%s", border.Bottom, border.Bottom)
			case yp == 0 && x == selectedPixel.x+1:
				pixStr = fmt.Sprintf("%s ", border.BottomRight)
			}
			bottom.WriteString(style.Render(pixStr))
		}
		if yp != len(bottomPixels)-1 {
			bottom.WriteRune('\n')
		}
	}
	if bottom.Len() > 0 {
		segments = append(segments, bottom.String())
	}

	charString := lipgloss.JoinVertical(lipgloss.Left, segments...)
	return SelectedCharStyle.Render(
		zone.Mark("selectedChar", charString),
	)
}

func (m FontViewer) ViewStatic() string {
	line := make([]string, 0, m.font.Width)
	lines := make([]string, 0, m.font.Height)
	for i, char := range m.font.Chars {
		if i != 0 && i%m.font.Width == 0 {
			lines = append(lines, lipgloss.JoinHorizontal(lipgloss.Top, line...))
			line = make([]string, 0, m.font.Width)
		}
		if i == m.selectedChar {
			line = append(line, viewSelectedChar(char, m.selectedPixel))
		} else {
			line = append(line, viewUnselectedChar(char, m.selectedPixel))
		}
		// Limit output so the screen is not overfilled
		if i == 79 {
			break
		}
	}
	lines = append(lines, lipgloss.JoinHorizontal(lipgloss.Top, line...))

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}

func (m FontViewer) View() string {
	return zone.Scan(m.ViewFollowChar())
}

func viewableCharsRange(char string, dim tea.WindowSizeMsg) (top, right, bottom, left int) {
	vertSpace := dim.Height/lipgloss.Height(char) - 1
	horiSpace := dim.Width/lipgloss.Width(char) - 1
	return vertSpace / 2, horiSpace/2 + horiSpace%2, vertSpace/2 + vertSpace%2, horiSpace / 2
}

func (m FontViewer) ViewFollowChar() string {
	selectedCharString := viewSelectedChar(m.font.Chars[m.selectedChar], m.selectedPixel)
	top, right, bottom, left := viewableCharsRange(selectedCharString, m.dimensions)
	// 4 chars to the left, selected char, 4 chars to the right
	line := make([]string, 0, left+1+right)
	// 2 chars above, selected char, 2 chars below
	lines := make([]string, 0, top+1+bottom)
	//lines = append(lines, fmt.Sprintf("%d", m.selectedChar))
	charX := m.selectedChar % m.font.Width
	charY := m.selectedChar / m.font.Width
	paddingLeft := max(left-charX, 0) * lipgloss.Width(selectedCharString)
	paddingTop := max(top-charY, 0) * (lipgloss.Height(selectedCharString))

	for y := max(charY-top, 0); y <= charY+bottom; y++ {
		if y == m.font.Height {
			// If it's the last line of the font add a single extra line, because otherwise the bottom borders of the chars is cut off
			// I don't know why it's needed, but it works
			lines = append(lines, "")
			break
		}
		for x := max(charX-left, 0); x <= charX+right && x < m.font.Width; x++ {
			if x == charX && y == charY {
				line = append(line, selectedCharString)
				continue
			}
			line = append(line, zone.Mark(fmt.Sprintf("char:x%d,y%d", x, y), viewUnselectedChar(m.font.Char(x, y), m.selectedPixel)))
		}
		lines = append(lines, lipgloss.JoinHorizontal(lipgloss.Top, line...))
		line = make([]string, 0, left+1+right)
	}

	return lipgloss.NewStyle().
		PaddingLeft(paddingLeft).
		PaddingTop(paddingTop).
		Render(lipgloss.JoinVertical(lipgloss.Left, lines...))
}
