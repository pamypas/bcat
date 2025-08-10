# bcat

**bcat** is a tiny Go command‑line utility that reads a Markdown (or any text) file or standard input, writes it to a temporary file, and opens it in a web browser (or any program that can handle `file://` URLs). It also supports a _tee_ mode to echo the raw data to stdout and allows selecting a custom browser command.

## Build

You can build the binary directly with Go:

```bash
go build -o bcat ./cmd/bcat
```

Or use the provided Makefile:

```bash
make        # builds all binaries into ./bin
make install # copies binaries to $(GOBIN) or $HOME/go/bin
```

## Usage

```bash
# Open a local Markdown file in the default browser
bcat path/to/file.md

# Pipe data from another command
cat README.md | bcat

# Write the data to stdout as well (tee)
cat README.md | bcat -tee

# Use a specific browser or command
bcat -browser "firefox" path/to/file.md
# or with stdin
cat README.md | bcat -browser "google-chrome"

# Tiny markdown viewer
# alias bcat='bcat -browser yandex-browser -tee | glow'
aichat --model openai:gpt-5 "Dog vs cat? Only answer" | bcat

  Dog
```

The `-browser` flag defaults to `xdg-open` on Linux, `open` on macOS, and `rundll32` on Windows, but you can override it with any command that accepts a URL as its first argument.
