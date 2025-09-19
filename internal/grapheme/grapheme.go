package grapheme

import (
	"regexp"
)

var reEol = regexp.MustCompile(`\r?\n`)
var reLetter = regexp.MustCompile(`\p{Letter}`)
var reSeparator = regexp.MustCompile(`\p{Separator}`)
var reOther = regexp.MustCompile(`\p{Other}`)

type Grapheme struct {
	seg         string
	Bytes       []byte
	Width       int
	IsLetter    bool
	IsSeparator bool
	IsOther     bool
	IsVisible   bool
	IsEol       bool
}

func NewGrapheme(seg string, bytes []byte, width int) *Grapheme {
	g := Grapheme{}

	g.seg = seg
	g.Bytes = bytes
	g.Width = width
	g.IsLetter = reLetter.MatchString(seg)
	g.IsSeparator = reSeparator.MatchString(seg)
	g.IsOther = reOther.MatchString(seg)
	g.IsVisible = !g.IsSeparator && !g.IsOther
	g.IsEol = reEol.MatchString(seg)

	return &g
}
