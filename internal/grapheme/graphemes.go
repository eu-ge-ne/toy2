package grapheme

var Graphemes *GraphemePool

func init() {
	m := map[string]*Grapheme{
		"\u0000": NewGrapheme("\u0000", []byte("␀"), 1),
		"\u0001": NewGrapheme("\u0001", []byte("␁"), 1),
		"\u0002": NewGrapheme("\u0002", []byte("␂"), 1),
		"\u0003": NewGrapheme("\u0003", []byte("␃"), 1),
		"\u0004": NewGrapheme("\u0004", []byte("␄"), 1),
		"\u0005": NewGrapheme("\u0005", []byte("␅"), 1),
		"\u0006": NewGrapheme("\u0006", []byte("␆"), 1),
		"\u0007": NewGrapheme("\u0007", []byte("␇"), 1),
		"\u0008": NewGrapheme("\u0008", []byte("␈"), 1),
		"\u0009": NewGrapheme("\u0009", []byte("\u2022\u2022\u2022\u2022"), 4), // "␉"
		"\u000a": NewGrapheme("\u000a", []byte("␊"), 1),
		"\u000b": NewGrapheme("\u000b", []byte("␋"), 1),
		"\u000c": NewGrapheme("\u000c", []byte("␌"), 1),
		"\u000d": NewGrapheme("\u000d", []byte("␍"), 1),
		"\u000e": NewGrapheme("\u000e", []byte("␎"), 1),
		"\u000f": NewGrapheme("\u000f", []byte("␏"), 1),
		"\u0010": NewGrapheme("\u0010", []byte("␐"), 1),
		"\u0011": NewGrapheme("\u0011", []byte("␑"), 1),
		"\u0012": NewGrapheme("\u0012", []byte("␒"), 1),
		"\u0013": NewGrapheme("\u0013", []byte("␓"), 1),
		"\u0014": NewGrapheme("\u0014", []byte("␔"), 1),
		"\u0015": NewGrapheme("\u0015", []byte("␕"), 1),
		"\u0016": NewGrapheme("\u0016", []byte("␖"), 1),
		"\u0017": NewGrapheme("\u0017", []byte("␗"), 1),
		"\u0018": NewGrapheme("\u0018", []byte("␘"), 1),
		"\u0019": NewGrapheme("\u0019", []byte("␙"), 1),
		"\u001a": NewGrapheme("\u001a", []byte("␚"), 1),
		"\u001b": NewGrapheme("\u001b", []byte("␛"), 1),
		"\u001c": NewGrapheme("\u001c", []byte("␜"), 1),
		"\u001d": NewGrapheme("\u001d", []byte("␝"), 1),
		"\u001e": NewGrapheme("\u001e", []byte("␞"), 1),
		"\u001f": NewGrapheme("\u001f", []byte("␟"), 1),
		"\u0020": NewGrapheme("\u0020", []byte("\u2027"), 1), // "␠"
		"\u007f": NewGrapheme("\u007f", []byte("␡"), 1),
		"\r\n":   NewGrapheme("\r\n", []byte("␍␊"), 2),
	}

	for i := 0x21; i < 0x7f; i += 1 {
		s := string(rune(i))

		m[s] = NewGrapheme(s, []byte(s), 1)
	}

	Graphemes = NewGraphemePool(m)
}
