package key

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
