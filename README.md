# GO Lang

## Introduction
- Compiler langugage
- Can run without any VMs
- Executables are different for OS
- OOPs: No concept of classes, there are structs, no overloading/overriding

## Packages
- We should define a go file with a package
- Why: to make the codes more organized and reachable
- We should have atleast one package in a project
- Packages are custom name, but entry point file has to be named as main
- We can have more than one packages in a paroject, we can use existing defined packages in a project, we can import the package in another file to use functions/variables.

## Lexer
- A lexer converts source code into clean, structured tokens so the parser can understand it.
- Significance:
    - Makes it easier to detect syntax errors.
    - Simplifies parsing.
    - Handles things like skipping spaces, reading numbers, detecting keywords.
    - First step in understanding or executing code.
- Lexer is a part of go compiler.
- How does a go compiler looks: ***lexer → parser → AST → code generation → binary***

## Types
- Case insensitive
- Variables to be declared in advance
- Everything is a type: like we say everythin is object in JS
- Different types:
    - string
    - bool
    - integer: uint8, uint64, int8, int64, uintptr
    - float: float32, float64
    - complex: iotas and all
    - Arrays
    - Slices
    - Maps
    - Struct
    - Pointers

## Commands
1. **To initialize a go project**: *go mod init folder-name*
    ```
    - Marks the dir as GO project
    - Creates go.mod: requirements file in case of go
    - It creates a module of the <folder-name>, as to build the code we have to define a module which can be run, a module then searches for a package main as entry point
    - To be performed outside the dir which you want to make a go project
    ```
2. **To Run**: *go run file_name.go*
3. **Help**: *go help*
4. **Go Path**: *go env GOPATH*
5. **Build**: *go build* to create an executable file for the project