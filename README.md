## What?
This project consists of a **compiler** (awoocc), **linker** (awoold), **simple object file dumper** (awoodump) and an **emulator** (awoomu) for the programming language **Awoo**.

**Awoo** is an **ease-of-use first** language similar with it's syntax to Go, making it easy to port existing Go code into the language.  
So far consisting only of **basic features**, the language lacks of complex implementation you would see in other languages.  

## How?
The **code** (.awoo) is compiled into separate **object files** (.awoobj) which are then linked together in a linker producing an **executable** (.awooxe).  
By default these steps are separeted into 3 directories, **src** for .awoo source code, **obj** for .awoobj object files and **bin** for .awooxe executables.  
This **executable** is further loaded into an emulator, which runs the code as instructed.  

Instruction set with which the code is compiled is **RISC-V** and it's extensions (even though my implementation is kinda scuffed).

## Why?
I wanted to learn some more about how things work low-level and besides having own compiler chain is kinda based so yeah.  

## Can I try?
You need **MinGW (64-bit version)** and a **Windows** computer (you're free to try on Linux, it should be the same lol).  
In addition, you need **SDL libraries** (basic ones and font) pasted into your **mingw64** directory.  
For starters, you might want to clone the [example repo](https://github.com/LamkasDev/awoo-example) (prints text on screen) into your documents directory under **.../Documents/awoo**.  
If you want to change location of the files, you're free to change **AWOODIR** variable inside **Makefile** to whatever you like.  
A **configuration file** named **emu.json** (under your **AWOODIR/config** directory) allows for setting clockrate of emulator. 

Lint: ```golangci-lint run```  
Unconvert: ```go run github.com/mdempsky/unconvert -tags="awoo32,awoodebug" -v ./cmd/...```

Only build (compiler): ```mingw32-make buildcc```  
Only build (linker): ```mingw32-make buildld```  
Only build (emulator): ```mingw32-make buildmu```  
Only build (object file dumper): ```mingw32-make builddump```  
Only build (all 4 tools): ```mingw32-make build```  

Build & Install (all 4 tools): ```mingw32-make install```  

Build & Run (compiler only): ```mingw32-make runcc```  
Build & Run (linker only): ```mingw32-make runld```  
Build & Run (emulator only): ```mingw32-make runmu```  
Build & Run (object file dumper): ```mingw32-make rundump```  
**Build & Run (all 3 required steps)**: ```mingw32-make run```

Clean: ```mingw32-make clean```  
