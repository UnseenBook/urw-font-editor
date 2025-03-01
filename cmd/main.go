package main

import (
	"fmt"
	"github.com/UnseenBook/urw-font-editor/editor"
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
	fmt.Println(font)
}
