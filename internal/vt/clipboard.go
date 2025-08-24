package vt

import (
	"encoding/base64"
	"fmt"
)

func copyToClipboard(text string) {
	write(osc(fmt.Sprintf("52;c;%s\x1b\\", base64.StdEncoding.EncodeToString([]byte(text)))))
}
