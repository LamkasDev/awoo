AWOOPLATFORM=awoo32
AWOODIR=$(USERPROFILE)\Documents\awoo

GOOS=windows
GOARCH=amd64
GOTAGS=$(AWOOPLATFORM),awoodebug

.PHONY: build install runll runemu run clean
build:
	@set GOOS=$(GOOS)
	@set GOARCH=$(GOARCH)
	@go build -o build/$(AWOOPLATFORM)/awoocc.exe -tags $(GOTAGS) cmd/awoocc/main.go
	@go build -o build/$(AWOOPLATFORM)/awoold.exe -tags $(GOTAGS) cmd/awoold/main.go
	@go build -o build/$(AWOOPLATFORM)/awoomu.exe -tags $(GOTAGS) cmd/awoomu/main.go

install: build
	@if exist "$(AWOODIR)\bin\$(AWOOPLATFORM)" rmdir /S /Q "$(AWOODIR)\bin\$(AWOOPLATFORM)"
	@xcopy "build\$(AWOOPLATFORM)" "$(AWOODIR)\bin\$(AWOOPLATFORM)\" /E /C /I >nul
	@if exist"$(AWOODIR)\resources" rmdir /S /Q "$(AWOODIR)\resources"
	@xcopy "resources" "$(AWOODIR)\resources\" /E /C /I >nul

runcc: build
	@if not exist "$(AWOODIR)\bin\dev" mkdir "$(AWOODIR)\bin\dev"
	@copy "build\$(AWOOPLATFORM)\awooll.exe" "$(AWOODIR)\bin\dev\awooll.exe" >nul
	@cd "build\$(AWOOPLATFORM)" && .\awooll.exe -i "$(AWOODIR)\data\input.awoo" -o "$(AWOODIR)\data\obj\input.awoobj"
	
runld: build
	@if not exist "$(AWOODIR)\bin\dev" mkdir "$(AWOODIR)\bin\dev"
	@copy "build\$(AWOOPLATFORM)\awoold.exe" "$(AWOODIR)\bin\dev\awoold.exe" >nul
	@cd "build\$(AWOOPLATFORM)" && .\awoold.exe -i "$(AWOODIR)\data\obj\input.awoobj" -o "$(AWOODIR)\data\bin\input.awooxe"

runemu: build
	@if not exist "$(AWOODIR)\bin\dev" mkdir "$(AWOODIR)\bin\dev"
	@copy "build\$(AWOOPLATFORM)\awoomu.exe" "$(AWOODIR)\bin\dev\awoomu.exe" >nul
	@cd "build\$(AWOOPLATFORM)" && .\awoomu.exe -i "$(AWOODIR)\data\bin\input.awooxe"

run: build
	@if exist "$(AWOODIR)\bin\dev" rmdir /S /Q "$(AWOODIR)\bin\dev"
	@xcopy "build\$(AWOOPLATFORM)" "$(AWOODIR)\bin\dev\" /E /C /I >nul
	@cd "build\$(AWOOPLATFORM)" && .\awoocc.exe -i "$(AWOODIR)\data\input.awoo" -o "$(AWOODIR)\data\obj\input.awoobj" -q
	@cd "build\$(AWOOPLATFORM)" && .\awoold.exe -i "$(AWOODIR)\data\obj\input.awoobj" -o "$(AWOODIR)\data\bin\input.awooxe" -q
	@cd "build\$(AWOOPLATFORM)" && .\awoomu.exe -i "$(AWOODIR)\data\bin\input.awooxe"

clean:
	@if exist "build" rmdir /S /Q build