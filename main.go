package main

import (
	"image/png"
	"log"
	"os"

	"github.com/UnseenBook/urw-font-editor/output"
)

func main() {
	const charCount, charWidth, charHeight = 256, 8, 8
	const imgHeight = (charHeight + 2) * charCount
	const imgWidth = charWidth

	outPutImage := output.NewOutputImage()

	font, err := os.Open("URW.FNT")
	if err != nil {
		log.Fatal(err)
	}
	defer font.Close()

	buffer := make([]byte, 80)

	for {
		count, err := font.Read(buffer)
		if count == 0 {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		_, err = outPutImage.Write(buffer)
		if err != nil {
			log.Fatal(err)
		}
	}

	outputFont, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFont.Close()

	if err := png.Encode(outputFont, outPutImage); err != nil {
		log.Fatal(err)
	}
}
