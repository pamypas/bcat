# Makefile for building the Go binaries in ./cmd
#
# This Makefile compiles all Go programs located in the ./cmd directory
# and installs them into the local bin directory.
#
# Usage:
#   make          # build all binaries
#   make install  # install binaries to $(GOBIN) or ./bin
#   make clean   # remove built binaries
#   make help    # show this help

# Directories
CMD_DIR   := cmd
BIN_DIR   := bin
GO        := go
GOFLAGS   := -v
GOBIN     := $(shell go env GOPATH)/bin

# Find all Go packages in ./cmd
PKGS := $(shell find $(CMD_DIR) -maxdepth 1 -type d -not -path "$(CMD_DIR)" -printf "%f\n")
# Build targets
TARGETS := $(addprefix $(BIN_DIR)/,$(PKGS))

.PHONY: all help install clean

all: $(TARGETS)

$(BIN_DIR)/%: $(CMD_DIR)/%/main.go
	@mkdir -p $(BIN_DIR)
	@echo "Building $@"
	$(GO) build $(GOFLAGS) -o $@ $<

install: $(TARGETS)
	@echo "Installing binaries to $(GOBIN)"
	@mkdir -p $(GOBIN)
	@for bin in $(TARGETS); do \
		echo "Installing $$bin to $(GOBIN)"; \
		cp $$bin $(GOBIN)/; \
	done

clean:
	@echo "Cleaning binaries"
	@rm -rf $(BIN_DIR)

help:
	@echo "Makefile commands:"
	@echo "  make          - build all binaries in $(CMD_DIR)"
	@echo "  make install  - install binaries to $(GOBIN)"
	@echo "  make clean   - remove built binaries"
	@echo "  make help    - show this help message"
