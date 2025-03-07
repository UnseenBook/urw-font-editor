package main

import (
	"fmt"
	"github.com/UnseenBook/urw-font-editor/editor"
	"github.com/UnseenBook/urw-font-editor/tui"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
)

func main() {
	inputFontFile, err := os.Open("URW.FNT")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFontFile.Close()
	fr := editor.NewUrwFontReader(inputFontFile)

	font, erro := fr.ReadFont()
	if erro != nil {
		log.Fatal(erro)
	}
	p := tea.NewProgram(
		tui.NewTui(font),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)
	if _, progErr := p.Run(); progErr != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}
