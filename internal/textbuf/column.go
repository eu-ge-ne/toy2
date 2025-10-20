package textbuf

func (buf *TextBuf) ColToByte(ln, col int) (int, bool) {
	i := 0

	for seg := range buf.LineSegments(ln) {
		if seg.Col == col {
			return i, true
		}

		i += len(seg.Gr.Str)
	}

	return 0, false
}
