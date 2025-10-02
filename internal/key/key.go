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
	KeyCode   rune
	ShiftCode rune
	BaseCode  rune
	Event     Event
	Text      string
	Mods      Mods
}

type Event uint8

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
		next_esc_i := bytes.IndexByte(raw[1:], 0x1b)
		if next_esc_i < 0 {
			next_esc_i = len(raw)
		} else {
			next_esc_i += 1
		}
		text := string(raw[:next_esc_i])
		return Key{Name: text, Text: text}, next_esc_i, true
	}

	match := re.FindSubmatch(raw)
	if match == nil {
		return Key{}, 0, false
	}

	key := Key{}

	if funcName, ok := funcNames[fmt.Sprintf("%s%s%s", match[1], match[2], match[8])]; ok {
		key.Name = funcName
	} else if i, err := strconv.Atoi(string(match[2])); err == nil {
		key.Name = string(rune(i))
	} else {
		key.Name = fmt.Sprintf("%s%s", match[1], match[8])
	}

	code, _ := strconv.Atoi(string(match[2]))
	shiftCode, _ := strconv.Atoi(string(match[3]))
	baseCode, _ := strconv.Atoi(string(match[4]))

	key.KeyCode = rune(code)
	key.ShiftCode = rune(shiftCode)
	key.BaseCode = rune(baseCode)

	mods, err := strconv.Atoi(string(match[5]))
	if err != nil {
		mods = 1
	}
	key.Mods = Mods(mods - 1)

	switch string(match[6]) {
	case "2":
		key.Event = EventRepeat
	case "3":
		key.Event = EventRelease
	default:
		key.Event = EventPress
	}

	if len(match[7]) > 0 {
		cps := strings.Split(string(match[7]), ":")
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
