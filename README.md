# README

A simple terminal command for copy and paste a content from device clipboard.

# Usage

Compiling
```
go build .
```

Copy a content from anywhere, using ctrl+C or using your terminal and fclip util
```
cat README.md | flip copy
```

Paste a content from clipboard to any destination, using fclip
```
fclip paste > README.md.new
```

## For developers
How to init go project
```
go mod init flip
go get github.com/atotto/clipboard
```
