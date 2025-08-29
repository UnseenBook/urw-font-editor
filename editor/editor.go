package editor

// Pixel is true when it is the foreground and false when it is the background
type Pixel bool

type Char [][]Pixel

type Font struct {
	Chars  []Char
	Width  int
	Height int
}

type FontReader interface {
	ReadFont() (Font, error)
}

func (f Font) Char(x, y int) Char {
	i := y*f.Width + x

	// Clamp i between 0 and the last index of Chars
	return f.Chars[min(max(i, 0), len(f.Chars)-1)]
}

func (f Font) FillFont(fr FontReader) (Font, error) {
	f, err := fr.ReadFont()
	if err != nil {
		return Font{}, err
	}
	return f, nil
}

func (f Font) TogglePixel(c, x, y int) Font {
	if c < 0 || c >= len(f.Chars) {
		return f
	}

	char := f.Chars[c]
	if x < 0 || x >= char.Width() || y < 0 || y >= char.Height() {
		return f
	}

	char[y][x] = !char[y][x]
	f.Chars[c] = char
	return f
}
