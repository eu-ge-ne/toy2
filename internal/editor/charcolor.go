package editor

type charColorEnum int

const (
	charColorUndefined charColorEnum = iota
	charColorVisible
	charColorWhitespace
	charColorEmpty
	charColorVisibleSelected
	charColorWhitespaceSelected
	charColorEmptySelected
)

func createCharColor(isSelected, isVisible, whitespaceEnabled bool) charColorEnum {
	if isSelected {
		if isVisible {
			return charColorVisibleSelected
		} else if whitespaceEnabled {
			return charColorWhitespaceSelected
		} else {
			return charColorEmptySelected
		}
	} else {
		if isVisible {
			return charColorVisible
		} else if whitespaceEnabled {
			return charColorWhitespace
		} else {
			return charColorEmpty
		}
	}
}
