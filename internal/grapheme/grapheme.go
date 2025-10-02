package grapheme

import (
	"regexp"
)

// var reLetter = regexp.MustCompile(`\p{Letter}`)
var reSeparator = regexp.MustCompile(`\p{Separator}`)
var reOther = regexp.MustCompile(`\p{Other}`)
var reEol = regexp.MustCompile(`\r?\n`)

type Grapheme struct {
	Seg   string
	Bytes []byte
	Width int
	//IsLetter    bool
	//IsSeparator bool
	//IsOther     bool
	IsVisible bool
	IsEol     bool
}

func NewGrapheme(seg string, bytes []byte, width int) *Grapheme {
	g := Grapheme{}

	g.Seg = seg
	g.Bytes = bytes
	g.Width = width
	//g.IsLetter = reLetter.MatchString(seg)
	isSeparator := reSeparator.MatchString(seg)
	isOther := reOther.MatchString(seg)
	g.IsVisible = !isSeparator && !isOther
	g.IsEol = reEol.MatchString(seg)

	return &g
}
