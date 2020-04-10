package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	const charCount, charWidth, charHeight = 256, 8, 8
	const imgHeight = (charHeight + 2) * charCount
	const imgWidth = charWidth

	img := image.NewGray(image.Rect(0, 0, charWidth, (charHeight+2)*charCount))

	file, err := os.Open("URW.FNT") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buffer := make([]byte, 1)

	for y := 0; y < imgHeight; y++ {
		for x := 0; x < imgWidth; x++ {
			_, err := file.Read(buffer)
			if err != nil {
				log.Fatal(err)
			}
			if buffer[0] == 0 {
				img.Set(x, y, color.Gray{255})
			} else {
				img.Set(x, y, color.Gray{0})
			}
		}
	}

	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
