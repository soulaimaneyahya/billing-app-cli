# Makefile

# Variables
SRC_DIR := ./src
DIST_DIR := ./dist
MAIN_BINARY := $(DIST_DIR)/main

# Default target
all: build

# Build target
build:
	mkdir -p $(DIST_DIR)
	go build -o $(MAIN_BINARY) $(SRC_DIR)

# Run target
run: build
	$(MAIN_BINARY)

# Clean target
clean:
	rm -f $(MAIN_BINARY)
