# Makefile

### ALMALINUX
# GOOS:= linux
# GOARCH:= amd64
# APP_NAME := RPS-Utilities-LINUX

### WINDOWS
GOOS:= windows
GOARCH:= amd64
APP_NAME := RPS-Utilities.exe

### RADXA
# GOOS:= linux
# GOARCH:= arm64
# APP_NAME := DVRStreamAdapter-Arm

###

VERSION := v3.0.0
BUILD_DATE := $(shell date +%Y-%m-%d_%H:%M:%S)
PKG := github.com/raulsilva-tech/rps-utilities
MAIN := main.go
OUTPUT := bin/$(APP_NAME)

CGO_ENABLED:= 0

build:
	@echo "Building $(APP_NAME)..."
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-s -w -X '$(PKG).APIVersion=$(VERSION)' -X '$(PKG).APIBuildDate=$(BUILD_DATE)'" -o $(OUTPUT) $(MAIN)

clean:
	@echo "Cleaning..."
	rm -f $(OUTPUT)

info:
	@echo "App:         $(APP_NAME)"
	@echo "Version:     $(VERSION)"
	@echo "Build date:  $(BUILD_DATE)"
	@echo "Main file:   $(MAIN)"
	@echo "Output:      $(OUTPUT)"

.PHONY: build clean info
