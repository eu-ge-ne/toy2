package editor

func (ed *Editor) HasChanges() bool {
	return !ed.history.IsEmpty()
}

func (ed *Editor) SetText(text string) {
	ed.buffer.Reset(text)
	ed.syntax.Reset()
}

func (ed *Editor) GetText() string {
	return ed.buffer.All()
}

func (ed *Editor) LoadFromFile(filePath string) error {
	err := ed.buffer.LoadFromFile(filePath)
	if err != nil {
		return err
	}

	ed.syntax.Reset()

	return nil
}

func (ed *Editor) SaveToFile(filePath string) error {
	return ed.buffer.SaveToFile(filePath)
}
