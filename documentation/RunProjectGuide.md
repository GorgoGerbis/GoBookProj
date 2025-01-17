# **Guide: Running Your Go Project**

## **Quick Reference**
1. **Run Project**: `go run main.go`
   - go run is used to execute regular Go files (like main.go) that contain a main() function.
2. **Build Binary**: `go build -o bin/monkey`
   - You run the go build -o bin/monkey command from the directory where your main.go file is located, which is the entry point of your application. 
   -  In your project structure, this would be: 'C:\Users\jacks\projects\writing_in_go_repo\GoBookProj\src\monkey'
3. **Run Tests**: `go test ./...`
   - go test is used to run *_test.go files that contain test functions starting with Test (e.g., func TestSomething(t *testing.T)).
4. **Install Dependencies**: `go mod tidy`
   - NOT NEEDED FOR THIS BOOK

## About Go:

The Go programming language is a compiled language, which means that your source code (written in `.go` files) is converted into a machine-readable binary executable before it can be run. Here’s a detailed explanation of how Go works and why you need to build your code:

### **How the Go Programming Language Works**
1. **Source Code**:
   - You write your program in `.go` files.
   - Example: `main.go` contains human-readable code written in Go syntax.

2. **Go Compiler**:
   - The `go build` command invokes the Go compiler.
   - The compiler translates your Go code into a **binary executable** that your operating system can run.
   - During this step:
     - Syntax and type errors are checked.
     - External dependencies are resolved and included in the binary.
     - The resulting binary is optimized for performance.

3. **Executable Binary**:
   - The output of the `go build` process is an executable file (e.g., `monkey` or `monkey.exe`).
   - This binary file contains machine code specific to your operating system and architecture (e.g., Windows, Linux, macOS).

4. **Execution**:
   - Once compiled, the binary can be run directly without needing the Go runtime or the original source code.
   - Example: `./bin/monkey` executes your program.

### **Go Development Workflow**
1. Write your source code in `.go` files.
2. Use `go build` to compile your program into a binary.
3. Run the binary directly (e.g., `./bin/monkey`).
4. For quick testing during development, you can skip building by using:
   ```bash
   go run main.go
   ```
   This combines the build and execute steps temporarily but does not create a standalone binary.

---

## **1. Understand Your Project Structure**
Based on your description, your project structure looks like this:
```
GoBookProj/
│
├── .envrc               # Environment-specific configurations
├── .gitignore           # Files to ignore in version control
├── README.md            # Project documentation
├── writing_an_interpreter_in_go_1.7.pdf
├── src/                 # Source code
│   └── monkey/
│       ├── go.mod       # Dependency management file
│       ├── main.go      # Entry point of your application
│       ├── lexer/       # Lexer package
│       │   ├── lexer.go
│       │   ├── lexer_test.go
│       └── token/       # Token package
│           └── token.go
```

## **2. Setup Go Environment**
Ensure you have Go installed on your system. Verify with:
```bash
go version
```
If not installed, download and install Go from [golang.org](https://golang.org).

## **3. Navigate to the Project Directory**
Go to your `src/monkey` directory where the `go.mod` and `main.go` files are located:
```bash
cd C:\Users\jacks\projects\writing_in_go_repo\GoBookProj\src\monkey
```

## **4. Install Dependencies** - NOT NEEDED FOR THIS BOOK
Run the following command to download and install any dependencies defined in `go.mod`:
```bash
go mod tidy
```
This ensures all required modules are installed and cleans up unnecessary entries.

## **5. Run the Project**
Execute the `main.go` file to run your application:
```bash
go run main.go
```
This will output "Hello, World!" or any updated functionality you have coded in `main.go`.

## **6. Build the Project**
You run the go build -o bin/monkey command from the directory where your main.go file is located, which is the entry point of your application. 
- In your project structure, this would be: 'C:\Users\jacks\projects\writing_in_go_repo\GoBookProj\src\monkey'

To create an executable binary:
```bash
go build -o bin/monkey
```
- The `-o bin/monkey` flag specifies the output location for the binary.
- This creates the `monkey` executable in the `bin/` directory.
- If the bin/ directory does not already exist, Go will create it for you.
- Ensure main.go is properly configured as the entry point. It must include the main() function.

Run the binary:
```bash
./bin/monkey
```

## **7. Test the Project**
Run the tests for the `lexer` package:
```bash
go test ./lexer
```
Run tests for the entire project:
```bash
go test ./...
```
This executes all `_test.go` files recursively.


## **8. Debugging**
Use the Go debugger to step through your code:
```bash
dlv debug main.go
```


