AWOOPLATFORM=awoo32
AWOODIR=$(USERPROFILE)\Documents\awoo

GOOS=windows
GOARCH=amd64
GOTAGS=$(AWOOPLATFORM),awoodebug

.PHONY: build install runll runemu run clean
build:
	@set GOOS=$(GOOS)
	@set GOARCH=$(GOARCH)
	@go build -o build/$(AWOOPLATFORM)/awoomu.exe -tags $(GOTAGS) cmd/awoomu/main.go
	@go build -o build/$(AWOOPLATFORM)/awooll.exe -tags $(GOTAGS) cmd/awooll/main.go

install: build
	@if not exist "$(AWOODIR)\bin\$(AWOOPLATFORM)" mkdir "$(AWOODIR)\bin\$(AWOOPLATFORM)"
	@copy "build\$(AWOOPLATFORM)\awoomu.exe" "$(AWOODIR)\bin\$(AWOOPLATFORM)\awoomu.exe" >nul
	@copy "build\$(AWOOPLATFORM)\awooll.exe" "$(AWOODIR)\bin\$(AWOOPLATFORM)\awooll.exe" >nul

runll: build
	@if not exist "$(AWOODIR)\bin\dev" mkdir "$(AWOODIR)\bin\dev"
	@copy "build\$(AWOOPLATFORM)\awooll.exe" "$(AWOODIR)\bin\dev\awooll.exe" >nul
	@./build/$(AWOOPLATFORM)/awooll.exe -i "$(AWOODIR)\data\input.awoo" -o obj

runemu: build
	@if not exist "$(AWOODIR)\bin\dev" mkdir "$(AWOODIR)\bin\dev"
	@copy "build\$(AWOOPLATFORM)\awoomu.exe" "$(AWOODIR)\bin\dev\awoomu.exe" >nul
	@./build/$(AWOOPLATFORM)/awoomu.exe -i "$(AWOODIR)\data\obj\input.awoobj"

run: build
	@if not exist "$(AWOODIR)\bin\dev" mkdir "$(AWOODIR)\bin\dev"
	@copy "build\$(AWOOPLATFORM)\awooll.exe" "$(AWOODIR)\bin\dev\awooll.exe" >nul
	@copy "build\$(AWOOPLATFORM)\awoomu.exe" "$(AWOODIR)\bin\dev\awoomu.exe" >nul
	@./build/$(AWOOPLATFORM)/awooll.exe -i "$(AWOODIR)\data\input.awoo" -o obj -q
	@./build/$(AWOOPLATFORM)/awoomu.exe -i "$(AWOODIR)\data\obj"

clean:
	@if exist "build" rmdir /S /Q build