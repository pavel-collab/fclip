# README

A simple terminal command for copy and paste a content from device clipboard.

# Usage

Compiling debug version
```
make debug
```

Compiling release version
```
make release
```

Run tests
```
make test
```

Copy a content from anywhere, using ctrl+C or using your terminal and fclip util
```
cat README.md | ./bin/flip copy
```

Paste a content from clipboard to any destination, using fclip
```
./bin/fclip paste > README.md.new
```