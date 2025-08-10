package main

import (
	"flag"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// openBrowser opens the given URL using the specified command.
// The command should be an executable that can accept a URL as its first argument.
func openBrowser(browserCmd string, url string) error {
	// If no command is provided, fall back to a sensible default based on the OS.
	if browserCmd == "" {
		switch runtime.GOOS {
		case "darwin":
			browserCmd = "open"
		case "windows":
			browserCmd = "rundll32"
		default: // linux, freebsd, etc.
			browserCmd = "xdg-open"
		}
		// For the Windows fallback we need to pass the additional argument.
		if runtime.GOOS == "windows" {
			return exec.Command(browserCmd, "url.dll,FileProtocolHandler", url).Start()
		}
	}
	// For non‑Windows commands we simply pass the URL as the first argument.
	return exec.Command(browserCmd, url).Start()
}

func main() {
	// Flags
	tee := flag.Bool("tee", false, "write output to stdout as well")
	browser := flag.String("browser", "xdg-open", "browser command to use (default: xdg-open)")
	flag.Parse()

	// Determine input source.
	var data []byte
	var err error
	if len(flag.Args()) == 0 {
		// No file argument: read from stdin.
		data, err = io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("Error reading stdin: %v", err)
		}
	} else {
		// First argument is a file path.
		filePath := flag.Args()[0]
		data, err = os.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Error reading file %s: %v", filePath, err)
		}
	}

	// If tee flag is set, write the raw data to stdout.
	if *tee {
		_, err := os.Stdout.Write(data)
		if err != nil {
			log.Fatalf("Error writing to stdout: %v", err)
		}
	}

	// Write data to a temporary file if it came from stdin.
	var targetPath string
	if len(flag.Args()) == 0 {
		tmpFile, err := os.CreateTemp("", "bcat-*.md")
		if err != nil {
			log.Fatalf("Error creating temp file: %v", err)
		}
		// defer os.Remove(tmpFile.Name())
		if _, err := tmpFile.Write(data); err != nil {
			log.Fatalf("Error writing to temp file: %v", err)
		}
		if err := tmpFile.Close(); err != nil {
			log.Fatalf("Error closing temp file: %v", err)
		}
		targetPath = tmpFile.Name()
	} else {
		targetPath = flag.Args()[0]
	}

	absPath, err := filepath.Abs(targetPath)
	if err != nil {
		log.Fatalf("Error getting absolute path: %v", err)
	}
	url := "file://" + absPath

	// Open the file in the selected browser.
	if err := openBrowser(*browser, url); err != nil {
		log.Fatalf("Failed to open browser: %v", err)
	}
}
