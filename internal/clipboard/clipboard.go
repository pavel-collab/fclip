package clipboard

import (
	"errors"
	"strings"

	"github.com/atotto/clipboard"
)

// ErrClipboardUnavailable is returned when the clipboard is unavailable.
var ErrClipboardUnavailable = errors.New("буфер обмена недоступен")

// ErrEmptyData is returned when attempting to copy empty data.
var ErrEmptyData = errors.New("данные для копирования пустые")

// Copy saves data to the clipboard.
// Returns an error if the data is empty or the clipboard is unavailable.
func Copy(data string) error {
    // Ensure the data is not empty
	if strings.TrimSpace(data) == "" {
		return ErrEmptyData
	}
	
    // Check clipboard availability
	if !clipboard.Unsupported {
		err := clipboard.WriteAll(data)
		if err != nil {
            // Check whether this is a clipboard unavailability error
			if strings.Contains(strings.ToLower(err.Error()), "no clipboard utilities available") {
				return ErrClipboardUnavailable
			}
			return err
		}
		return nil
	}
	
	return ErrClipboardUnavailable
}

// Paste retrieves data from the clipboard.
// Returns an error if the clipboard is unavailable.
func Paste() (string, error) {
    // Check clipboard availability
	if !clipboard.Unsupported {
		data, err := clipboard.ReadAll()
		if err != nil {
            // Check whether this is a clipboard unavailability error
			if strings.Contains(strings.ToLower(err.Error()), "no clipboard utilities available") {
				return "", ErrClipboardUnavailable
			}
			return "", err
		}
		return data, nil
	}
	
	return "", ErrClipboardUnavailable
}

// IsAvailable checks whether the clipboard is available on the current system.
func IsAvailable() bool {
	return !clipboard.Unsupported
}

// GetStatus returns the clipboard status.
func GetStatus() string {
	if IsAvailable() {
		return "доступен"
	}
	return "недоступен"
}
