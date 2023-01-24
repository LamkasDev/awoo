AWOOPLATFORM=awoo32

GOOS=windows
GOARCH=amd64
GOTAGS=$(AWOOPLATFORM),awoodebug

.PHONY: build run clean
build:
	@set GOOS=$(GOOS)
	@set GOARCH=$(GOARCH)
	@go build -o build/$(AWOOPLATFORM)/awoomu.exe -tags $(GOTAGS) cmd/awoomu/main.go
	@go build -o build/$(AWOOPLATFORM)/awooll.exe -tags $(GOTAGS) cmd/awooll/main.go

runemu: build
	@./build/$(AWOOPLATFORM)/awoomu.exe

runll: build
	@./build/$(AWOOPLATFORM)/awooll.exe

clean:
	@if exist "build" rmdir /S /Q build