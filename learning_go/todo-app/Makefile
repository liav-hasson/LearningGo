BINARY_NAME=todo

# go can handle linux forward slash on windows
CMD_DIR=./cmd

# Add Go to PATH for MSYS2/Git Bash where it's not in PATH by default
export PATH := /c/Program Files/Go/bin:$(PATH)

.PHONY: windows linux run clean

# Build for Windows (adds .exe automatically)
windows:
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME).exe $(CMD_DIR)
	@echo "Created Windows executable!"

linux:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux $(CMD_DIR)
	@echo "Created Linux executable!"

build-all: windows linux

run:
	go run $(CMD_DIR)

# Clean up binaries. wont work on powershell, use on git-bash instead
clean:
	rm -f $(BINARY_NAME).exe $(BINARY_NAME)-linux