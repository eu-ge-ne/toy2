package syntax

type highlightReq struct {
	ln0 int
	ln1 int
	hl chan *Highlighter
}
