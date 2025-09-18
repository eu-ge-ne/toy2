package vt

import (
	"encoding/base64"
	"fmt"
)

func copyToClipboard(text string) {
	Sync.Write(osc(fmt.Sprintf("52;c;%s\x1b\\", base64.StdEncoding.EncodeToString([]byte(text)))))
}
