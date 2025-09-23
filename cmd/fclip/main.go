package main

import (
	"fmt"
	"io"
	"os"

	"fclip/internal/clipboard"
)

// main - точка входа для CLI утилиты.
// Поддерживает команды: copy (чтение из stdin в буфер обмена), paste (вывод в stdout).
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Использование:")
		fmt.Println("  fclip copy   # копировать из stdin в буфер обмена")
		fmt.Println("  fclip paste  # вставить из буфера обмена в stdout")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "copy":
		// Чтение из stdin
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка чтения:", err)
			os.Exit(1)
		}
		if err := clipboard.Copy(string(data)); err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка записи в буфер обмена:", err)
			os.Exit(1)
		}

	case "paste":
		// Вывод содержимого буфера обмена
		data, err := clipboard.Paste()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка чтения из буфера обмена:", err)
			os.Exit(1)
		}
		fmt.Print(data)

	default:
		fmt.Fprintln(os.Stderr, "Неизвестная команда:", os.Args[1])
		os.Exit(1)
	}
}
