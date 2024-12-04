To set up a Go programming project, you need to organize your files and follow standard practices to make your project maintainable and compatible with Go tools. Here's a step-by-step guide:

---

### 1. **Install Go**
Ensure Go is installed on your system. You can download it from [golang.org](https://golang.org).

### 2. **Set Up Your Workspace**
Decide on a folder where you’ll store your Go projects. For example:
```bash
mkdir ~/go-projects
cd ~/go-projects
```

### 3. **Initialize Your Project**
Use `go mod init` to create a new Go module. A Go module is a collection of Go packages stored in a directory with a `go.mod` file at its root.

```bash
mkdir my-go-project
cd my-go-project
go mod init github.com/yourusername/my-go-project
```

Replace `github.com/yourusername/my-go-project` with the module path that matches where your project will be hosted (e.g., GitHub, GitLab, etc.).

### 4. **Typical File Structure**
Here’s a common structure for a Go project:

```
my-go-project/
├── cmd/                # Entry points (e.g., main packages for CLI tools)
│   └── app/            # Example app entry point
│       └── main.go
├── pkg/                # Reusable packages
│   └── example/        # Example package
│       ├── example.go
│       └── example_test.go
├── internal/           # Private packages (not imported outside this module)
│   └── helpers/
│       └── helper.go
├── go.mod              # Module definition
├── go.sum              # Dependency checksums
└── README.md           # Documentation
```

---

### 5. **Create Your `main.go`**
Start with an entry point for your application:
```bash
mkdir -p cmd/app
touch cmd/app/main.go
```

```go
// cmd/app/main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

---

### 6. **Add Additional Packages**
Create reusable or internal packages as needed:
```bash
mkdir -p pkg/example
touch pkg/example/example.go
```

```go
// pkg/example/example.go
package example

func Hello() string {
    return "Hello from example package!"
}
```

Use the package in `main.go`:
```go
// cmd/app/main.go
package main

import (
    "fmt"
    "github.com/yourusername/my-go-project/pkg/example"
)

func main() {
    fmt.Println(example.Hello())
}
```

---

### 7. **Run Your Application**
To run your application, use:
```bash
go run cmd/app/main.go
```

---

### 8. **Testing**
Go uses the `testing` package for unit tests. Create test files alongside your Go files, named `*_test.go`.

Example:
```bash
touch pkg/example/example_test.go
```

```go
// pkg/example/example_test.go
package example

import "testing"

func TestHello(t *testing.T) {
    expected := "Hello from example package!"
    if Hello() != expected {
        t.Errorf("Expected %s, got %s", expected, Hello())
    }
}
```

Run tests with:
```bash
go test ./...
```

---

### 9. **Manage Dependencies**
Add dependencies using `go get`:
```bash
go get github.com/some/dependency
```

Dependencies are automatically tracked in `go.mod` and checksums are added to `go.sum`.

---

### 10. **Build Your Application**
To compile your application into an executable:
```bash
go build -o bin/app cmd/app/main.go
```

Run the compiled binary:
```bash
./bin/app
```

---

### 11. **Version Control**
Create a `.gitignore` file:
```
bin/
*.exe
*.log
*.tmp
```

Initialize Git and push to your repository:
```bash
git init
git add .
git commit -m "Initial commit"
git remote add origin <repository-url>
git push -u origin main
```

---

### Summary
- **Initialize the module:** `go mod init`.
- **Organize files:** Follow the typical structure with `cmd/`, `pkg/`, and `internal/`.
- **Build and run:** Use `go build` and `go run`.
- **Test your code:** Add `*_test.go` files for unit tests.
- **Manage dependencies:** Use `go mod` and `go get`.

This structure and workflow will make your Go project modular, testable, and easy to maintain.