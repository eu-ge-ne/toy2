package syntax

type editKind int

const (
	editDelete editKind = iota
	editInsert
)

type edit struct {
	edit     editKind
	startLn  int
	startCol int
	endLn    int
	endCol   int
}
