package app

import (
	"io"
	"os"
	"unicode/utf8"
)

func (app *App) load() error {
	f, err := os.Open(app.filePath)
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

		app.editor.Buffer.Append(string(chunk))
	}

	return nil
}
