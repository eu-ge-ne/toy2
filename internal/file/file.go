package file

import (
	"io"
	"os"
	"unicode/utf8"

	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

func Load(filePath string, textBuf *textbuf.TextBuf) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer f.Close()

	buf := make([]byte, 1024*1024*64)

	for {
		bytesRead, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		chunk := buf[:bytesRead]

		if !utf8.Valid(chunk) {
			panic("invalid utf8 chunk")
		}

		textBuf.Append(string(chunk))
	}

	return nil
}
