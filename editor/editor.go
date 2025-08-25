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
