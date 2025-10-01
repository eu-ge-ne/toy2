package key

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Key struct {
	Name      string
	KeyCode   int
	ShiftCode int
	BaseCode  int
	Event     Event
	Text      string
	Shift     bool
	Alt       bool
	Ctrl      bool
	Super     bool
	CapsLock  bool
	NumLock   bool
}

type Event int

const (
	EventPress Event = iota
	EventRepeat
	EventRelease
)

const rePrefix = `(\x1b\x5b|\x1b\x4f)`
const reCodes = `(?:(\d+)(?::(\d*))?(?::(\d*))?)?`
const reParams = `(?:;(\d*)?(?::(\d*))?)?`
const reCodepoints = `(?:;([\d:]*))?`
const reScheme = `([u~ABCDEFHPQS])`

var re = regexp.MustCompile(rePrefix + reCodes + reParams + reCodepoints + reScheme)

func Parse(raw []byte) (Key, int, bool) {
	if len(raw) == 0 {
		return Key{}, 0, false
	}

	b := raw[0]

	switch {
	case b == 0x1b && len(raw) == 1:
		return Key{Name: "ESC"}, 1, true
	case b == 0x0d:
		return Key{Name: "ENTER"}, 1, true
	case b == 0x09:
		return Key{Name: "TAB"}, 1, true
	case b == 0x7f || b == 0x08:
		return Key{Name: "BACKSPACE"}, 1, true
	case b != 0x1b:
		next_esc_i := bytes.Index(raw[1:], []byte{0x1b})
		if next_esc_i < 0 {
			next_esc_i = len(raw)
		} else {
			next_esc_i += 1
		}
		text := string(raw[:next_esc_i])
		return Key{Name: text, Text: text}, next_esc_i, true
	}

	match := re.FindStringSubmatch(string(raw))
	if match == nil {
		return Key{}, 0, false
	}

	key := Key{}

	if funcName, ok := funcNames[fmt.Sprintf("%s%s%s", match[1], match[2], match[8])]; ok {
		key.Name = funcName
	} else if i, err := strconv.Atoi(match[2]); err == nil {
		key.Name = string(rune(i))
	} else {
		key.Name = fmt.Sprintf("%s%s", match[1], match[8])
	}

	code, _ := strconv.Atoi(match[2])
	shiftCode, _ := strconv.Atoi(match[3])
	baseCode, _ := strconv.Atoi(match[4])

	key.KeyCode = code
	key.ShiftCode = shiftCode
	key.BaseCode = baseCode

	mods, err := strconv.Atoi(match[5])
	if err != nil {
		mods = 1
	}
	mods -= 1

	key.Shift = mods&1 == 1
	key.Alt = mods&2 == 2
	key.Ctrl = mods&4 == 4
	key.Super = mods&8 == 8
	key.CapsLock = mods&64 == 64
	key.NumLock = mods&128 == 128

	switch match[6] {
	case "2":
		key.Event = EventRepeat
	case "3":
		key.Event = EventRelease
	default:
		key.Event = EventPress
	}

	if len(match[7]) > 0 {
		cps := strings.Split(match[7], ":")
		runes := make([]rune, 0, len(cps))

		for _, cp := range cps {
			i, _ := strconv.Atoi(cp)
			runes = append(runes, rune(i))
		}

		key.Text = string(runes)
	}

	index := re.FindIndex(raw)

	return key, index[1] - index[0], true
}

var funcNames = map[string]string{
	"\x1b[27u":  "ESC",
	"\x1b[13u":  "ENTER",
	"\x1b[9u":   "TAB",
	"\x1b[127u": "BACKSPACE",

	"\x1b[2~": "INSERT",
	"\x1b[3~": "DELETE",

	"\x1b[1D": "LEFT",
	"\x1b[D":  "LEFT",
	"\x1bOD":  "LEFT",
	"\x1b[1C": "RIGHT",
	"\x1b[C":  "RIGHT",
	"\x1bOC":  "RIGHT",
	"\x1b[1A": "UP",
	"\x1b[A":  "UP",
	"\x1bOA":  "UP",
	"\x1b[1B": "DOWN",
	"\x1b[B":  "DOWN",
	"\x1bOB":  "DOWN",

	"\x1b[5~": "PAGE_UP",
	"\x1b[6~": "PAGE_DOWN",

	"\x1b[7~": "HOME",
	"\x1b[1H": "HOME",
	"\x1b[H":  "HOME",
	"\x1bOH":  "HOME",
	"\x1b[8~": "END",
	"\x1b[1F": "END",
	"\x1b[F":  "END",
	"\x1bOF":  "END",

	"\x1b[57358u": "CAPS_LOCK",
	"\x1b[57359u": "SCROLL_LOCK",
	"\x1b[57360u": "NUM_LOCK",
	"\x1b[57361u": "PRINT_SCREEN",
	"\x1b[57362u": "PAUSE",
	"\x1b[57363u": "MENU",

	"\x1b[11~": "F1",
	"\x1b[1P":  "F1",
	"\x1b[P":   "F1",
	"\x1bOP":   "F1",

	"\x1b[12~": "F2",
	"\x1b[1Q":  "F2",
	"\x1b[Q":   "F2",
	"\x1bOQ":   "F2",

	"\x1b[13~": "F3",
	"\x1bOR":   "F3",

	"\x1b[14~": "F4",
	"\x1b[1S":  "F4",
	"\x1b[S":   "F4",
	"\x1bOS":   "F4",

	"\x1b[15~": "F5",
	"\x1b[17~": "F6",
	"\x1b[18~": "F7",
	"\x1b[19~": "F8",
	"\x1b[20~": "F9",
	"\x1b[21~": "F10",
	"\x1b[23~": "F11",
	"\x1b[24~": "F12",

	"\x1b[57441u": "LEFT_SHIFT",
	"\x1b[57442u": "LEFT_CONTROL",
	"\x1b[57443u": "LEFT_ALT",
	"\x1b[57444u": "LEFT_SUPER",
	"\x1b[57447u": "RIGHT_SHIFT",
	"\x1b[57448u": "RIGHT_CONTROL",
	"\x1b[57449u": "RIGHT_ALT",
	"\x1b[57450u": "RIGHT_SUPER",
}
