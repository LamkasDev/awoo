AWOOPLATFORM=awoo64

GOOS=windows
GOARCH=amd64
GOTAGS=$(AWOOPLATFORM),awoodebug

.PHONY: build run clean
build:
	@set GOOS=$(GOOS)
	@set GOARCH=$(GOARCH)
	@go build -o build/$(AWOOPLATFORM)/awoo-emu.exe -tags $(GOTAGS) cmd/main.go

run: build
	@./build/$(AWOOPLATFORM)/awoo-emu.exe

clean:
	@rmdir /S /Q build