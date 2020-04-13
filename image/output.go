package image

import (
	"fmt"
	"image"
	"image/color"
)

// NewOutputImage creates a new output image
func NewOutputImage() *Image {
	return &Image{image.NewGray(image.Rect(0, 0, imgWidth, imgHeight)), 0}
}

func (i *Image) Write(buffer []byte) (count int, err error) {
	if len(buffer) != charWidth*charHeight {
		err = fmt.Errorf("write buffer should be the size of one character e.g. %d", charWidth*charHeight)
		return
	}

	x0, y0, err := i.nextCharPos()
	if err != nil {
		return
	}
	i.charCount++

	for count = 0; count < len(buffer); count++ {
		x := (count % charWidth) + x0
		y := count/charWidth + y0
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
