package editor

// Pixel is true when it is the foreground and false when it is the background
type Pixel bool

type Char [][]Pixel

type Font struct {
	Chars []Char
}

type FontReader interface {
	ReadFont() (Font, error)
}

func (f Font) FillFont(fr FontReader) (Font, error) {
	f, err := fr.ReadFont()
	if err != nil {
		return Font{}, err
	}
	return f, nil
}
