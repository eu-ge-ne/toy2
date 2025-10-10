package syntax

type editKind int

const (
	editReset editKind = iota
	editDelete
	editInsert
)

type edit struct {
	edit     editKind
	startLn  int
	startCol int
	endLn    int
	endCol   int
}
