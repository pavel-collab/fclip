package clipboard

import (
	"errors"
	"strings"

	"github.com/atotto/clipboard"
)

// ErrClipboardUnavailable возвращается, когда буфер обмена недоступен.
var ErrClipboardUnavailable = errors.New("буфер обмена недоступен")

// ErrEmptyData возвращается при попытке копирования пустых данных.
var ErrEmptyData = errors.New("данные для копирования пустые")

// Copy сохраняет данные в буфер обмена.
// Возвращает ошибку, если данные пустые или буфер обмена недоступен.
func Copy(data string) error {
	// Проверяем, что данные не пустые
	if strings.TrimSpace(data) == "" {
		return ErrEmptyData
	}
	
	// Проверяем доступность буфера обмена
	if !clipboard.Unsupported {
		err := clipboard.WriteAll(data)
		if err != nil {
			// Проверяем, является ли это ошибкой недоступности буфера обмена
			if strings.Contains(strings.ToLower(err.Error()), "no clipboard utilities available") {
				return ErrClipboardUnavailable
			}
			return err
		}
		return nil
	}
	
	return ErrClipboardUnavailable
}

// Paste извлекает данные из буфера обмена.
// Возвращает ошибку, если буфер обмена недоступен.
func Paste() (string, error) {
	// Проверяем доступность буфера обмена
	if !clipboard.Unsupported {
		data, err := clipboard.ReadAll()
		if err != nil {
			// Проверяем, является ли это ошибкой недоступности буфера обмена
			if strings.Contains(strings.ToLower(err.Error()), "no clipboard utilities available") {
				return "", ErrClipboardUnavailable
			}
			return "", err
		}
		return data, nil
	}
	
	return "", ErrClipboardUnavailable
}

// IsAvailable проверяет, доступен ли буфер обмена в текущей системе.
func IsAvailable() bool {
	return !clipboard.Unsupported
}

// GetStatus возвращает статус буфера обмена.
func GetStatus() string {
	if IsAvailable() {
		return "доступен"
	}
	return "недоступен"
}
