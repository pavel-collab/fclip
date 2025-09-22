package main

import (
 "fmt"
 "io"
 "os"

 "github.com/atotto/clipboard"
)

func main() {
 if len(os.Args) < 2 {
  fmt.Println("Использование:")
  fmt.Println("  myclip copy   # копировать из stdin в буфер обмена")
  fmt.Println("  myclip paste  # вставить из буфера обмена в stdout")
  os.Exit(1)
 }

 switch os.Args[1] {
 case "copy":
  // читаем данные из stdin и сохраняем в буфер обмена
  data, err := io.ReadAll(os.Stdin)
  if err != nil {
   fmt.Fprintln(os.Stderr, "Ошибка чтения:", err)
   os.Exit(1)
  }
  if err := clipboard.WriteAll(string(data)); err != nil {
   fmt.Fprintln(os.Stderr, "Ошибка записи в буфер обмена:", err)
   os.Exit(1)
  }

 case "paste":
  // выводим данные из буфера обмена
  data, err := clipboard.ReadAll()
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
