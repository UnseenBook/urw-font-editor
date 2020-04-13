package image

import (
	"fmt"
	"image"
)

const maxChars, charWidth, charHeight = 256, 8, 10

// CharSize is the amount of bytes per character
const CharSize = charWidth * charHeight
const charsPerRow = 16
const imgHeight = charHeight * maxChars / charsPerRow
const imgWidth = charWidth * charsPerRow

// Image is the output image and holds a counter for the character count
type Image struct {
	*image.Gray
	charCount int
}

func (i *Image) nextCharPos() (x int, y int, err error) {
	if i.charCount >= maxChars {
		err = fmt.Errorf("max character count of %d reached", maxChars)
		return
	}

	x = (i.charCount % charsPerRow) * charWidth
	y = (i.charCount / charsPerRow) * charHeight

	return
}
