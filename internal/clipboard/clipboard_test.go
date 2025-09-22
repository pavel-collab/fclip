package clipboard

import (
	"strings"
	"testing"
)

// TestCopyPasteBasic проверяет простое копирование и вставку.
func TestCopyPasteBasic(t *testing.T) {
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

// TestCopyEmpty проверяет вставку пустой строки.
func TestCopyEmpty(t *testing.T) {
	input := ""
	if err := Copy(input); err != nil {
		t.Fatalf("ошибка Copy: %v", err)
	}
	output, err := Paste()
	if err != nil {
		t.Fatalf("ошибка Paste: %v", err)
	}
	if output != input {
		t.Errorf("ожидалось пустое значение, получено %q", output)
	}
}

// TestCopyMultiline проверяет копирование многострочного текста.
func TestCopyMultiline(t *testing.T) {
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
	_, err := Paste()
	if err != nil {
		t.Errorf("Paste вернул ошибку: %v", err)
	}
}

// TestOverwrite проверяет, что новое значение перезаписывает старое.
func TestOverwrite(t *testing.T) {
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
