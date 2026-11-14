package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	GITHUB_TOKEN     = "GITHUB_ACCESS_TOKEN"
	GITHUB_ENTRY     = "https://github.com"
	GITHUB_API_ENTRY = "https://api.github.com"
	REPO             = "helix-editor/helix"
	ARCH             = "x86_64-linux"
	INSTALL_BIN      = "/usr/local/bin/hx"
	RUNTIME_DIR      = "$HOME/.config/helix/runtime"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load .env")
	}

	github_token := os.Getenv(GITHUB_TOKEN)
	if github_token == "" {
		panic(fmt.Sprintf("missing env(%s) from .env", GITHUB_TOKEN))
	}
}
