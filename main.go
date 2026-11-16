package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

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

func getLatestRelease() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	ctx := context.Background()

	apiURL, err := url.JoinPath(GITHUB_API_ENTRY, "repos", REPO, "releases", "latest")
	if err != nil {
		panic(err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)

	req.Header.Add("Accept", "application/vnd.github+json")

	res, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	fmt.Println("Status:", res.Status)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Body:", string(body))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load .env")
	}

	github_token := os.Getenv(GITHUB_TOKEN)
	if github_token == "" {
		panic(fmt.Sprintf("missing env(%s) from .env", GITHUB_TOKEN))
	}

	getLatestRelease()
}
