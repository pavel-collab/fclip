package clipboard

import (
	"strings"
	"testing"
	
	"github.com/atotto/clipboard"
)

// TestCopyPasteBasic проверяет простое копирование и вставку.
func TestCopyPasteBasic(t *testing.T) {
	// Пропускаем тест, если буфер обмена недоступен
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

// TestCopyEmpty проверяет обработку пустой строки.
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

// TestCopyMultiline проверяет копирование многострочного текста.
func TestCopyMultiline(t *testing.T) {
	// Пропускаем тест, если буфер обмена недоступен
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

// TestCopyUnicode проверяет работу с unicode-символами.
func TestCopyUnicode(t *testing.T) {
	// Пропускаем тест, если буфер обмена недоступен
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

// TestPasteWithoutCopy проверяет поведение, если перед этим не делали Copy.
// ⚠️ Здесь нельзя гарантировать пустое значение, т.к. в буфере может быть
// что-то из других программ. Мы лишь проверяем, что Paste не падает.
func TestPasteWithoutCopy(t *testing.T) {
	// Пропускаем тест, если буфер обмена недоступен
	if !IsAvailable() {
		t.Skip("буфер обмена недоступен в данной среде")
	}
	
	_, err := Paste()
	if err != nil {
		t.Errorf("Paste вернул ошибку: %v", err)
	}
}

// TestOverwrite проверяет, что новое значение перезаписывает старое.
func TestOverwrite(t *testing.T) {
	// Пропускаем тест, если буфер обмена недоступен
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

// TestLargeInput проверяет работу с большим текстом.
func TestLargeInput(t *testing.T) {
	// Пропускаем тест, если буфер обмена недоступен
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

// TestClipboardUnavailable проверяет поведение при недоступном буфере обмена.
func TestClipboardUnavailable(t *testing.T) {
	if IsAvailable() {
		t.Skip("буфер обмена доступен, тест пропущен")
	}
	
	// Тестируем Copy с недоступным буфером обмена
	err := Copy("test")
	if err == nil {
		t.Fatal("ожидалась ошибка при копировании с недоступным буфером обмена")
	}
	if err != ErrClipboardUnavailable {
		t.Errorf("ожидалась ошибка ErrClipboardUnavailable, получено: %v", err)
	}
	
	// Тестируем Paste с недоступным буфером обмена
	_, err = Paste()
	if err == nil {
		t.Fatal("ожидалась ошибка при вставке с недоступным буфером обмена")
	}
	if err != ErrClipboardUnavailable {
		t.Errorf("ожидалась ошибка ErrClipboardUnavailable, получено: %v", err)
	}
}

// TestIsAvailable проверяет функцию проверки доступности буфера обмена.
func TestIsAvailable(t *testing.T) {
	// Функция должна всегда возвращать валидное значение
	available := IsAvailable()
	if available != !clipboard.Unsupported {
		t.Errorf("IsAvailable() вернула %v, ожидалось %v", available, !clipboard.Unsupported)
	}
}

// TestGetStatus проверяет функцию получения статуса буфера обмена.
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

// TestCopyWhitespaceOnly проверяет обработку строк, содержащих только пробелы.
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
