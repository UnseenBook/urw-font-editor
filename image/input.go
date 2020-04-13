package image

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

// NewInputImage creates a new output image
func NewInputImage() *Image {
	reader, err := os.Open("input.png")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	imageFile, err := png.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	imageGray, ok := imageFile.(*image.Gray)
	if !ok {
		log.Fatalf("wrong color scheme for input image, actual: %T", imageFile.ColorModel())
	}

	return &Image{imageGray, 0}
}

func (i *Image) Read(buffer []byte) (count int, err error) {
	if len(buffer) != CharSize {
		err = fmt.Errorf("can only read with buffer size %d", CharSize)
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
		buffer[count] = i.mapGrayToByte(i.At(x, y))
	}

	return
}

func (i *Image) mapGrayToByte(anyColor color.Color) byte {
	gray := color.GrayModel.Convert(anyColor).(color.Gray)
	if gray.Y == 255 {
		return byte(0)
	}
	return byte(1)
}
