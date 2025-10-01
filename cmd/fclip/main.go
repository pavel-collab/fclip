package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"fclip/internal/clipboard"
)

const (
	exitSuccess = 0
	exitError   = 1
)

// main - entry point for the CLI utility.
// Supports commands: copy (read from stdin to clipboard), paste (print to stdout).
func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(exitError)
	}

	command := strings.ToLower(strings.TrimSpace(os.Args[1]))
	
	switch command {
	case "copy":
		if err := handleCopy(); err != nil {
        fmt.Fprintf(os.Stderr, "Ошибка при копировании: %v\n", err)
			os.Exit(exitError)
		}
	case "paste":
		if err := handlePaste(); err != nil {
        fmt.Fprintf(os.Stderr, "Ошибка при вставке: %v\n", err)
			os.Exit(exitError)
		}
	case "status":
		handleStatus()
		os.Exit(exitSuccess)
	case "help", "-h", "--help":
		printUsage()
		os.Exit(exitSuccess)
	default:
        fmt.Fprintf(os.Stderr, "Неизвестная команда: %s\n", command)
        fmt.Fprintln(os.Stderr, "Используйте 'fclip help' для просмотра доступных команд")
		os.Exit(exitError)
	}
}

// printUsage prints usage help for the program.
func printUsage() {
	fmt.Println("fclip - утилита для работы с буфером обмена")
	fmt.Println()
	fmt.Println("Использование:")
	fmt.Println("  fclip copy   # копировать из stdin в буфер обмена")
	fmt.Println("  fclip paste  # вставить из буфера обмена в stdout")
	fmt.Println("  fclip status # проверить статус буфера обмена")
	fmt.Println("  fclip help   # показать эту справку")
	fmt.Println()
	fmt.Println("Примеры:")
	fmt.Println("  echo 'Hello' | fclip copy")
	fmt.Println("  fclip paste > file.txt")
}

// handleCopy handles the copy command.
func handleCopy() error {
    // Ensure stdin is not empty
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("ошибка чтения из stdin: %w", err)
	}
	
    // Ensure the data is not empty
	dataStr := string(data)
	if strings.TrimSpace(dataStr) == "" {
		return fmt.Errorf("stdin пустой или содержит только пробельные символы")
	}
	
    // Validate data size (limit for safety)
	const maxSize = 1024 * 1024 // 1MB
	if len(data) > maxSize {
		return fmt.Errorf("данные слишком большие (максимум %d байт)", maxSize)
	}
	
	if err := clipboard.Copy(dataStr); err != nil {
		return fmt.Errorf("ошибка записи в буфер обмена: %w", err)
	}
	
	return nil
}

// handlePaste handles the paste command.
func handlePaste() error {
	data, err := clipboard.Paste()
	if err != nil {
		return fmt.Errorf("ошибка чтения из буфера обмена: %w", err)
	}
	
	fmt.Print(data)
	return nil
}

// handleStatus handles the command that checks clipboard status.
func handleStatus() {
	if clipboard.IsAvailable() {
		fmt.Println("✅ Буфер обмена доступен")
	} else {
		fmt.Println("❌ Буфер обмена недоступен")
		fmt.Println("💡 Установите xsel, xclip, wl-clipboard или Termux:API add-on")
	}
}
