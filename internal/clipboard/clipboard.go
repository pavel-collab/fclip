package clipboard

import (
	"github.com/atotto/clipboard"
)

// Copy сохраняет данные в буфер обмена.
func Copy(data string) error {
	return clipboard.WriteAll(data)
}

// Paste извлекает данные из буфера обмена.
func Paste() (string, error) {
	return clipboard.ReadAll()
}
