package editor

import (
	"fmt"
	"io"
)

type UrwFontReader struct {
	file io.Reader
}

func NewUrwFontReader(file io.Reader) UrwFontReader {
	return UrwFontReader{file: file}
}

func (f UrwFontReader) ReadFont() (Font, error) {
	chars := make([]Char, 0, 16*16)

	for i := 0; i < 16*16; i++ {
		c, err := f.readChar()
		if err != nil {
			return Font{}, err
		}
		chars = append(chars, c)
	}

	return Font{Chars: chars}, nil
}

func (f UrwFontReader) readChar() (Char, error) {
	byteCount := 8 * 10
	buf := make([]byte, byteCount)
	n, err := f.file.Read(buf)
	if err != nil {
		return Char{}, err
	}
	if n != byteCount {
		return Char{}, fmt.Errorf("expected file byte count as a multiple of 80, but ended with %d bytes", n)
	}

	char := make([][]Pixel, 0, 10)
	line := make([]Pixel, 0, 8)
	for i, b := range buf {
		if i != 0 && i%8 == 0 {
			char = append(char, line)
			line = make([]Pixel, 0, 8)
		}
		p := Pixel(b == 1)
		line = append(line, p)
	}
	char = append(char, line)

	return char, nil
}
