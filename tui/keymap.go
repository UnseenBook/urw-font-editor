package tui

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Up,
	Down,
	Left,
	Right,
	CharUp,
	CharDown,
	CharLeft,
	CharRight,
	TogglePixel,
	Quit key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "down"),
		),
		Left: key.NewBinding(
			key.WithKeys("left", "h"),
			key.WithHelp("←/h", "left"),
		),
		Right: key.NewBinding(
			key.WithKeys("right", "l"),
			key.WithHelp("→/l", "right"),
		),
		CharUp: key.NewBinding(
			key.WithKeys("ctrl+up", "ctrl+k"),
			key.WithHelp("ctrl+up/ctrl+k", "Char up"),
		),
		CharDown: key.NewBinding(
			key.WithKeys("ctrl+down", "ctrl+j"),
			key.WithHelp("ctrl+down/ctrl+j", "Char down"),
		),
		CharLeft: key.NewBinding(
			key.WithKeys("ctrl+left", "ctrl+h"),
			key.WithHelp("ctrl+left/ctrl+h", "Char left"),
		),
		CharRight: key.NewBinding(
			key.WithKeys("ctrl+right", "ctrl+l"),
			key.WithHelp("ctrl+right/ctrl+l", "Char right"),
		),
		TogglePixel: key.NewBinding(
			key.WithKeys(" "),
			key.WithHelp("space", "Toggle pixel"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q/ctrl+c", "Quit"),
		),
	}
}
