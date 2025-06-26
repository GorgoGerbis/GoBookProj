# Monkey Language Interpreter

A complete interpreter for the Monkey programming language, built from scratch in Go. This project follows along with **"Writing an Interpreter in Go"** by Thorsten Ball, which teaches the fundamentals of creating an interpreter from the ground up.

**Note**: This project is currently in development as I work through the book's concepts. I plan to extend it with additional features and optimizations once the core foundation is solid.

## What is Monkey?

Monkey is a simple but powerful programming language that supports:
- Variables and basic arithmetic
- Functions and closures  
- Arrays and hash maps
- Built-in functions
- Interactive REPL

```monkey
let fibonacci = fn(x) {
  if (x == 0) {
    0
  } else {
    if (x == 1) {
      1
    } else {
      fibonacci(x - 1) + fibonacci(x - 2);
    }
  }
};

fibonacci(10); // => 55
```

See more examples in the [`examples/`](examples/) directory.

## Architecture

The interpreter follows a clean, modular design:

- **Lexer** (`src/monkey/lexer/`) - tokenizes source code into meaningful symbols
- **Parser** (`src/monkey/parser/`) - builds an Abstract Syntax Tree using recursive descent
- **AST** (`src/monkey/ast/`) - tree structures representing program syntax  
- **REPL** (`src/monkey/repl/`) - interactive shell for testing and exploration

## Quick Start

```bash
# clone the repository
git clone https://github.com/GorgoGerbis/GoBookProj.git
cd GoBookProj/src/monkey

# run the REPL
go run main.go

# run tests  
go test ./...
```

### Running the REPL

```
Hello! This is the Monkey programming language!
Feel free to type in commands
>> let x = 5;
>> x + 10
15
>> let add = fn(a, b) { a + b };
>> add(5, 5)
10
```

## Current Status

✅ **Complete**
- Lexical analysis with full token support
- Recursive descent parser for expressions and statements
- AST node structures for all language constructs  
- Interactive REPL with basic functionality
- Comprehensive test coverage

🚧 **In Progress**  
- Expression evaluation engine
- Variable binding and scoping
- Function definitions and calls

## Testing

The project includes comprehensive test suites:

```bash
# run all tests
go test ./...

# test with coverage
go test -cover ./...

# run specific component tests
go test ./lexer
go test ./parser
```

## Language Features (Planned)

- **Data Types**: integers, booleans, strings, arrays, hash maps
- **Operators**: arithmetic (`+`, `-`, `*`, `/`), comparison (`==`, `!=`, `<`, `>`), logical (`!`)
- **Control Flow**: if/else expressions
- **Functions**: first-class functions with closures
- **Built-ins**: `len()`, `first()`, `last()`, `rest()`, `push()`

## Implementation Highlights

- **Pratt parser** for handling operator precedence elegantly
- **Test-driven development** with comprehensive edge case coverage
- **Clean separation** between lexing, parsing, and evaluation phases
- **Idiomatic Go** following standard conventions and best practices

## Learning Journey

This project represents my journey through fundamental computer science concepts:
- **Lexical Analysis**: Breaking source code into meaningful tokens
- **Recursive Descent Parsing**: Building syntax trees from token streams  
- **AST Evaluation**: Interpreting tree structures to execute programs
- **Language Design**: Understanding how programming languages work under the hood

## Credit

Built while following **"Writing an Interpreter in Go"** by Thorsten Ball - an excellent resource for learning interpreter implementation. The book provides clear explanations of complex concepts and guides you through building a complete interpreter step by step.

## Future Plans

Once I complete the book's core implementation, I plan to extend the interpreter with:
- Additional data structures (sets, maps with more features)
- Enhanced error reporting and debugging tools
- Performance optimizations and benchmarking
- Potential compilation to bytecode for faster execution
