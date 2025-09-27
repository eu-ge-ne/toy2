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

func Save(filePath string, textBuf *textbuf.TextBuf) error {
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	for text := range textBuf.Read(0) {
		_, err := f.WriteString(text)
		if err != nil {
			return err
		}
	}

	return nil
}
