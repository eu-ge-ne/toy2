package key

type Event uint8

const (
	Press Event = iota
	Repeat
	Release
)

func parseEvent(b []byte) Event {
	switch string(b) {
	case "2":
		return Repeat
	case "3":
		return Release
	}
	return Press
}
