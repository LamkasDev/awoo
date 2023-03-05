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
	@if exist "$(AWOODIR)\bin\$(AWOOPLATFORM)" rmdir /S /Q "$(AWOODIR)\bin\$(AWOOPLATFORM)"
	@xcopy "build\$(AWOOPLATFORM)" "$(AWOODIR)\bin\$(AWOOPLATFORM)\" /E /C /I >nul
	@if exist"$(AWOODIR)\resources" rmdir /S /Q "$(AWOODIR)\resources"
	@xcopy "resources" "$(AWOODIR)\resources\" /E /C /I >nul

runll: build
	@if not exist "$(AWOODIR)\bin\dev" mkdir "$(AWOODIR)\bin\dev"
	@copy "build\$(AWOOPLATFORM)\awooll.exe" "$(AWOODIR)\bin\dev\awooll.exe" >nul
	@cd "build\$(AWOOPLATFORM)" && .\awooll.exe -i "$(AWOODIR)\data\input.awoo" -o obj

runemu: build
	@if not exist "$(AWOODIR)\bin\dev" mkdir "$(AWOODIR)\bin\dev"
	@copy "build\$(AWOOPLATFORM)\awoomu.exe" "$(AWOODIR)\bin\dev\awoomu.exe" >nul
	@cd "build\$(AWOOPLATFORM)" && .\awoomu.exe -i "$(AWOODIR)\data\obj\input.awoobj"

run: build
	@if exist "$(AWOODIR)\bin\dev" rmdir /S /Q "$(AWOODIR)\bin\dev"
	@xcopy "build\$(AWOOPLATFORM)" "$(AWOODIR)\bin\dev\" /E /C /I >nul
	@cd "build\$(AWOOPLATFORM)" && .\awooll.exe -i "$(AWOODIR)\data\input.awoo" -o obj -q
	@cd "build\$(AWOOPLATFORM)" && .\awoomu.exe -i "$(AWOODIR)\data\obj"

clean:
	@if exist "build" rmdir /S /Q build