package output

import (
	"fmt"
	"image"
	"image/color"
)

const maxChars, charWidth, charHeight = 256, 8, 10
const charsPerRow = 1
const imgHeight = charHeight * maxChars / charsPerRow
const imgWidth = charWidth * charsPerRow

type Image struct {
	*image.Gray
	charCount int
}

func NewOutputImage() *Image {
	return &Image{image.NewGray(image.Rect(0, 0, imgWidth, imgHeight)), 0}
}

func (i *Image) nextCharPos() (x int, y int, err error) {
	if i.charCount >= maxChars {
		err = fmt.Errorf("Max character count of %d reached.", maxChars)
		return
	}

	x = (i.charCount * charWidth) % charsPerRow
	y = i.charCount * charHeight / charsPerRow

	return
}

func (i *Image) Write(buffer []byte) (count int, err error) {
	if len(buffer) != charWidth*charHeight {
		err = fmt.Errorf("Write buffer should be the size of one character e.g. %d", charWidth*charHeight)
		return
	}

	x0, y0, err := i.nextCharPos()
	if err != nil {
		return
	}
	i.charCount++

	for count = 0; count < len(buffer); count++ {
		x := (x0 + count) % charWidth
		y := (y0 + count) / charHeight
		i.SetGray(x, y, i.mapByteToGray(buffer[count]))
	}

	return
}

func (i *Image) mapByteToGray(code byte) color.Gray {
	if code == byte(0) {
		return color.Gray{255}
	}
	return color.Gray{0}
}
