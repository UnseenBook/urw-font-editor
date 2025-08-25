package main

import (
	"fmt"
	"log"
	"os"

	"github.com/UnseenBook/urw-font-editor/editor"
	"github.com/UnseenBook/urw-font-editor/tui"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func main() {
	inputFontFile, err := os.Open("URW.FNT")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFontFile.Close()
	fr := editor.NewUrwFontReader(inputFontFile)

	zone.NewGlobal()
	defer zone.Close()
	font, erro := fr.ReadFont()
	if erro != nil {
		log.Fatal(erro)
	}
	p := tea.NewProgram(
		tui.NewFontViewer(font),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)
	if _, progErr := p.Run(); progErr != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}
