package clipboard

import (
	"strings"
	"testing"
	
	"github.com/atotto/clipboard"
)

// TestCopyPasteBasic verifies simple copy and paste.
func TestCopyPasteBasic(t *testing.T) {
    // Skip the test if the clipboard is unavailable
	if !IsAvailable() {
        t.Skip("буфер обмена недоступен в данной среде")
	}
	
	input := "Hello, Clipboard!"
	if err := Copy(input); err != nil {
        t.Fatalf("ошибка Copy: %v", err)
	}
	output, err := Paste()
	if err != nil {
        t.Fatalf("ошибка Paste: %v", err)
	}
	if output != input {
        t.Errorf("ожидалось %q, получено %q", input, output)
	}
}

// TestCopyEmpty verifies handling of an empty string.
func TestCopyEmpty(t *testing.T) {
	input := ""
	err := Copy(input)
	if err == nil {
        t.Fatal("ожидалась ошибка при копировании пустой строки")
	}
	if err != ErrEmptyData {
        t.Errorf("ожидалась ошибка ErrEmptyData, получено: %v", err)
	}
}

// TestCopyMultiline verifies copying of multiline text.
func TestCopyMultiline(t *testing.T) {
    // Skip the test if the clipboard is unavailable
	if !IsAvailable() {
        t.Skip("буфер обмена недоступен в данной среде")
	}
	
	input := "Первая строка\nВторая строка\nТретья строка"
	if err := Copy(input); err != nil {
        t.Fatalf("ошибка Copy: %v", err)
	}
	output, err := Paste()
	if err != nil {
        t.Fatalf("ошибка Paste: %v", err)
	}
	if output != input {
        t.Errorf("ожидалось %q, получено %q", input, output)
	}
}

// TestCopyUnicode verifies handling of unicode characters.
func TestCopyUnicode(t *testing.T) {
    // Skip the test if the clipboard is unavailable
	if !IsAvailable() {
        t.Skip("буфер обмена недоступен в данной среде")
	}
	
	input := "Привет 🌍 こんにちは 你好"
	if err := Copy(input); err != nil {
        t.Fatalf("ошибка Copy: %v", err)
	}
	output, err := Paste()
	if err != nil {
        t.Fatalf("ошибка Paste: %v", err)
	}
	if output != input {
        t.Errorf("ожидалось %q, получено %q", input, output)
	}
}

// TestPasteWithoutCopy verifies behavior if Copy was not called before.
// ⚠️ We cannot guarantee an empty value here since the clipboard may contain
// data from other programs. We only verify that Paste does not fail.
func TestPasteWithoutCopy(t *testing.T) {
    // Skip the test if the clipboard is unavailable
	if !IsAvailable() {
        t.Skip("буфер обмена недоступен в данной среде")
	}
	
	_, err := Paste()
	if err != nil {
        t.Errorf("Paste вернул ошибку: %v", err)
	}
}

// TestOverwrite verifies that the new value overwrites the old one.
func TestOverwrite(t *testing.T) {
    // Skip the test if the clipboard is unavailable
	if !IsAvailable() {
        t.Skip("буфер обмена недоступен в данной среде")
	}
	
	first := "Первое значение"
	second := "Второе значение"

	if err := Copy(first); err != nil {
        t.Fatalf("ошибка Copy первого значения: %v", err)
	}
	if err := Copy(second); err != nil {
        t.Fatalf("ошибка Copy второго значения: %v", err)
	}
	output, err := Paste()
	if err != nil {
        t.Fatalf("ошибка Paste: %v", err)
	}
	if output != second {
        t.Errorf("ожидалось %q, получено %q", second, output)
	}
}

// TestLargeInput verifies handling of large text.
func TestLargeInput(t *testing.T) {
    // Skip the test if the clipboard is unavailable
	if !IsAvailable() {
        t.Skip("буфер обмена недоступен в данной среде")
	}
	
	var builder strings.Builder
	for i := 0; i < 10000; i++ {
		builder.WriteString("строка теста\n")
	}
	input := builder.String()

	if err := Copy(input); err != nil {
        t.Fatalf("ошибка Copy: %v", err)
	}
	output, err := Paste()
	if err != nil {
        t.Fatalf("ошибка Paste: %v", err)
	}
	if output != input {
        t.Errorf("ожидалось совпадение больших данных, но строки различаются")
	}
}

// TestClipboardUnavailable verifies behavior when the clipboard is unavailable.
func TestClipboardUnavailable(t *testing.T) {
	if IsAvailable() {
        t.Skip("буфер обмена доступен, тест пропущен")
	}
	
    // Test Copy with unavailable clipboard
	err := Copy("test")
	if err == nil {
        t.Fatal("ожидалась ошибка при копировании с недоступным буфером обмена")
	}
	if err != ErrClipboardUnavailable {
        t.Errorf("ожидалась ошибка ErrClipboardUnavailable, получено: %v", err)
	}
	
    // Test Paste with unavailable clipboard
	_, err = Paste()
	if err == nil {
        t.Fatal("ожидалась ошибка при вставке с недоступным буфером обмена")
	}
	if err != ErrClipboardUnavailable {
        t.Errorf("ожидалась ошибка ErrClipboardUnavailable, получено: %v", err)
	}
}

// TestIsAvailable verifies the function that checks clipboard availability.
func TestIsAvailable(t *testing.T) {
    // The function should always return a valid value
	available := IsAvailable()
	if available != !clipboard.Unsupported {
        t.Errorf("IsAvailable() вернула %v, ожидалось %v", available, !clipboard.Unsupported)
	}
}

// TestGetStatus verifies the function that returns the clipboard status.
func TestGetStatus(t *testing.T) {
	status := GetStatus()
	expected := "доступен"
	if !IsAvailable() {
		expected = "недоступен"
	}
	if status != expected {
        t.Errorf("GetStatus() вернула %q, ожидалось %q", status, expected)
	}
}

// TestCopyWhitespaceOnly verifies handling of strings that contain only whitespace.
func TestCopyWhitespaceOnly(t *testing.T) {
	testCases := []string{
		"   ",
		"\t\t\t",
		"\n\n\n",
		" \t\n ",
		"\r\n\r\n",
	}
	
	for _, input := range testCases {
		t.Run("input_"+strings.Replace(strings.Replace(strings.Replace(strings.Replace(input, " ", "s", -1), "\t", "t", -1), "\n", "n", -1), "\r", "r", -1), func(t *testing.T) {
			err := Copy(input)
			if err == nil {
                t.Fatal("ожидалась ошибка при копировании строки только с пробелами")
			}
			if err != ErrEmptyData {
                t.Errorf("ожидалась ошибка ErrEmptyData, получено: %v", err)
			}
		})
	}
}
