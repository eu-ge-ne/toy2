package vt

import (
	"encoding/base64"
	"fmt"
	"io"
)

func CopyToClipboard(out io.Writer, text string) {
	b64 := base64.StdEncoding.EncodeToString([]byte(text))

	fmt.Fprintf(out, "\x1b]52;c;%s\x1b\\", b64)
}
