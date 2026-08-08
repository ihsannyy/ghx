# Detect Operating System and Environment
UNAME_S := $(shell uname -s 2>/dev/null || echo Windows)

# Detect Termux vs Standard Linux/macOS/Windows install directory
ifneq ($(wildcard /data/data/com.termux/files/usr/bin),)
    PREFIX ?= /data/data/com.termux/files/usr
    INSTALL_DIR ?= $(PREFIX)/bin
else ifeq ($(UNAME_S),Darwin)
    INSTALL_DIR ?= /usr/local/bin
else ifeq ($(findstring NT,$(UNAME_S)),NT)
    INSTALL_DIR ?= $(USERPROFILE)/bin
else
    INSTALL_DIR ?= /usr/local/bin
endif

.PHONY: all deps build build-all install clean

all: deps build

deps:
	@echo "Checking required system dependencies (go, gh, git)..."
	@which go >/dev/null 2>&1 || ( \
		echo "Go not found. Installing..." && \
		if command -v pkg >/dev/null 2>&1; then pkg install -y golang; \
		elif command -v brew >/dev/null 2>&1; then brew install go; \
		elif command -v apt-get >/dev/null 2>&1; then sudo apt-get update && sudo apt-get install -y golang; \
		elif command -v dnf >/dev/null 2>&1; then sudo dnf install -y golang; \
		elif command -v pacman >/dev/null 2>&1; then sudo pacman -S --noconfirm go; \
		else echo "Error: Could not auto-install Go. Please install Go manually."; exit 1; fi \
	)
	@which gh >/dev/null 2>&1 || ( \
		echo "GitHub CLI (gh) not found. Installing..." && \
		if command -v pkg >/dev/null 2>&1; then pkg install -y gh; \
		elif command -v brew >/dev/null 2>&1; then brew install gh; \
		elif command -v apt-get >/dev/null 2>&1; then sudo apt-get install -y gh; \
		elif command -v dnf >/dev/null 2>&1; then sudo dnf install -y gh; \
		elif command -v pacman >/dev/null 2>&1; then sudo pacman -S --noconfirm github-cli; \
		else echo "Error: Could not auto-install gh CLI. Please install gh manually."; exit 1; fi \
	)
	@which git >/dev/null 2>&1 || ( \
		echo "Git not found. Installing..." && \
		if command -v pkg >/dev/null 2>&1; then pkg install -y git; \
		elif command -v brew >/dev/null 2>&1; then brew install git; \
		elif command -v apt-get >/dev/null 2>&1; then sudo apt-get install -y git; \
		elif command -v dnf >/dev/null 2>&1; then sudo dnf install -y git; \
		elif command -v pacman >/dev/null 2>&1; then sudo pacman -S --noconfirm git; \
		else echo "Error: Could not auto-install git. Please install git manually."; exit 1; fi \
	)

build:
	CGO_ENABLED=0 go build -o bin/ghx .

build-all:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bin/ghx-linux-arm64 .
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/ghx-linux-amd64 .
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o bin/ghx-darwin-arm64 .
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/ghx-darwin-amd64 .
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/ghx-windows-amd64.exe .

install: deps build
	@mkdir -p $(INSTALL_DIR)
	cp bin/ghx $(INSTALL_DIR)/ghx
	chmod +x $(INSTALL_DIR)/ghx
	@echo "✔ ghx successfully installed to $(INSTALL_DIR)/ghx"

clean:
	rm -rf bin
