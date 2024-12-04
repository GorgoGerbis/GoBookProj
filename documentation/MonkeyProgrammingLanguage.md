## The Monkey Programming Language & Interpreter
Every interpreter is built to interpret a specific programming language. That’s how you “implement”
a programming language. Without a compiler or an interpreter a programming language
is nothing more than an idea or a specification.

We’re going to parse and evaluate our own language called **Monkey**. It’s a language specifically
designed for this book. Its only implementation is the one we’re going to build in this book -
our interpreter.
### Expressed as a list of features, Monkey has the following:
• C-like syntax
• variable bindings
• integers and booleans
• arithmetic expressions
• built-in functions
• first-class and higher-order functions
• closures
• a string data structure
• an array data structure
• a hash data structure
### Variable Bindings
```go
let age = 1;
let name = "Monkey";
let result = 10 * (20 / 2);
````

Besides integers, booleans, and strings, the Monkey interpreter we’re going to build will also support arrays and hashes.

### Arrays

Here’s what binding an array of integers to a name looks like:

```go
let myArray = [1, 2, 3, 4, 5];
```

### Hashes

And here is a hash, where values are associated with keys:

```go
let thorsten = {"name": "Thorsten", "age": 28};
```

### Accessing Elements

Accessing the elements in arrays and hashes is done with index expressions:

```go
myArray[0];        // => 1
thorsten["name"];  // => "Thorsten"
```

### Functions

The `let` statements can also be used to bind functions to names. Here’s a small function that adds two numbers:

```go
let add = fn(a, b) { return a + b; };
```

Monkey supports **implicit return values**, which means we can leave out the `return` if we want to:

```go
let add = fn(a, b) { a + b; };
```

Calling a function is as easy as you’d expect:

```go
add(1, 2);
```

### Recursive Functions

A more complex function, such as a Fibonacci function that returns the Nth Fibonacci number, might look like this:

```go
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
```

Note the recursive calls to `fibonacci` itself!

### Higher-Order Functions

Monkey also supports a special type of functions called **higher-order functions**, which take other functions as arguments. Here’s an example:

```go
let twice = fn(f, x) {
  return f(f(x));
};
let addTwo = fn(x) {
  return x + 2;
};
twice(addTwo, 2); // => 6
```