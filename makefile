# Makefile for compiling mongpt

# Set the Go compiler to use
GO := go

# Set the flags to use when compiling
FLAGS := -v

# Set the output directory for the compiled binary
OUT_DIR := ./bin

# Set the name of the compiled binary
OUT_NAME := mongpt

# Default target
.PHONY: all
all: clean build

# Target for cleaning up previous builds
.PHONY: clean
clean:
	rm -rf $(OUT_DIR)/$(OUT_NAME)

# Target for building the binary
.PHONY: build
build: clean
	$(GO) build $(FLAGS) -o $(OUT_DIR)/$(OUT_NAME) main.go
	chmod +x $(OUT_DIR)/$(OUT_NAME)
	
