package main

import (
	"image/png"
	"log"
	"os"

	"github.com/UnseenBook/urw-font-editor/image"
)

func main() {
	inputImage := image.NewInputImage()

	fontBytes := make([]byte, 0, 2560)
	buffer := make([]byte, image.CharSize)
	for {
		count, err := inputImage.Read(buffer)
		if count == 0 {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fontBytes = append(fontBytes, buffer...)
	}

	outputFont, err := os.Create("New_URW.FNT")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFont.Close()
	_, err = outputFont.Write(fontBytes)
	if err != nil {
		log.Fatal(err)
	}

	outputImage := image.NewOutputImage()

	font, err := os.Open("URW.FNT")
	if err != nil {
		log.Fatal(err)
	}
	defer font.Close()

	buffer = make([]byte, image.CharSize)

	for {
		count, err := font.Read(buffer)
		if count == 0 {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		_, err = outputImage.Write(buffer)
		if err != nil {
			log.Fatal(err)
		}
	}

	outputImageFile, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}
	defer outputImageFile.Close()

	if err := png.Encode(outputImageFile, outputImage); err != nil {
		log.Fatal(err)
	}
}
