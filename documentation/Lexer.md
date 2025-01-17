### **Notable Data Structures in Your Project**

#### **1. `Token` (Defined in `token.go`)**
   - **Purpose**: Represents the smallest units of the Monkey language.
   - **Fields**:
     - `Type` (`TokenType`): Specifies the kind of token (e.g., `LET`, `IDENT`, `INT`).
     - `Literal` (`string`): The actual value of the token (e.g., `"5"`, `"let"`).
   - **Usage**: The lexer generates tokens to represent source code for further parsing.

#### **2. `Lexer` (Defined in `lexer.go`)**
   - **Purpose**: Processes input source code character by character and emits tokens.
   - **Fields**:
     - `input` (`string`): Holds the source code being processed.
     - `position` (`int`): Tracks the current character's position in `input`.
     - `readPosition` (`int`): Tracks the next character's position in `input` (used for peeking).
     - `ch` (`byte`): Stores the current character being examined.
   - **Usage**: Iteratively calls `NextToken()` to convert input into a stream of tokens.

#### **3. `keywords` Map (Defined in `token.go`)**
   - **Purpose**: Maps strings to their corresponding `TokenType` for language keywords like `let`, `fn`, etc.
   - **Type**: `map[string]TokenType`.
   - **Usage**: Helps distinguish user-defined identifiers from reserved keywords.

#### **4. `struct` for Test Cases (Defined in `lexer_test.go`)**
   - **Purpose**: Represents test cases for verifying lexer output.
   - **Fields**:
     - `expectedType` (`TokenType`): Expected token type.
     - `expectedLiteral` (`string`): Expected token literal.
   - **Usage**: Used to validate the lexer's tokenization logic.

#### **5. REPL Loop (Defined in `repl.go`)**
   - **Purpose**: Continuously reads user input, tokenizes it using the lexer, and outputs tokens.
   - **Components**:
     - `Scanner` (`bufio.Scanner`): Reads user input line by line.
     - `io.Reader` and `io.Writer`: Interfaces for handling input/output streams.

---

### **Function Call Walkthrough: Input "let x = 5 + 5;"**

#### **1. Program Entry Point**
   - `main.go`:
     - The `main` function starts and initializes the REPL by calling:
       ```go
       repl.Start(os.Stdin, os.Stdout)
       ```

#### **2. REPL Loop**
   - `repl.go`:
     - The `Start` function begins a loop:
       1. Prompts the user with `>> `.
       2. Reads a line of input using `Scanner.Text()`.
       3. Initializes a `Lexer` with the input:
          ```go
          l := lexer.New(line)
          ```

#### **3. Token Generation**
   - `lexer.go`:
     - The `Lexer` processes the input step by step:
       - **Initialization**:
         - The `New` function sets up the `Lexer` and calls `readChar` to load the first character.
       - **Tokenization Loop**:
         - Calls `NextToken()` repeatedly until `EOF`.
         - Each call:
           1. Skips whitespace (`skipWhitespace`).
           2. Matches the current character to a known token type using a `switch` statement.
           3. Advances the `position` and `readPosition` pointers to prepare for the next token.

#### **4. Token Emission**
   - Example for `"let x = 5 + 5;"`:
     - **Iteration 1**: Processes `"let"`:
       - Calls `isLetter` to recognize the keyword.
       - Emits `{Type: "LET", Literal: "let"}`.
     - **Iteration 2**: Processes `"x"`:
       - Calls `isLetter` for the identifier.
       - Emits `{Type: "IDENT", Literal: "x"}`.
     - **Iteration 3**: Processes `"="`:
       - Emits `{Type: "ASSIGN", Literal: "="}`.
     - **Iteration 4**: Processes `"5"`:
       - Calls `isDigit` for the number.
       - Emits `{Type: "INT", Literal: "5"}`.
     - **Iteration 5**: Processes `"+"`:
       - Emits `{Type: "PLUS", Literal: "+"}`.
     - **Iteration 6**: Processes `"5"`:
       - Emits `{Type: "INT", Literal: "5"}`.
     - **Iteration 7**: Processes `";"`:
       - Emits `{Type: "SEMICOLON", Literal: ";"}`.

#### **5. REPL Output**
   - Each token is printed to the output:
     ```plaintext
     {Type: LET, Literal: let}
     {Type: IDENT, Literal: x}
     {Type: ASSIGN, Literal: =}
     {Type: INT, Literal: 5}
     {Type: PLUS, Literal: +}
     {Type: INT, Literal: 5}
     {Type: SEMICOLON, Literal: ;}
     ```

---

The **lookup table** in your project is implemented as the `keywords` map in `token.go`. Here's a detailed explanation of its role and importance:

---

### **Lookup Table: `keywords` Map**

#### **Purpose**
- The lookup table (`keywords` map) helps differentiate **identifiers** (e.g., variable names like `x`, `y`) from **reserved keywords** (e.g., `let`, `fn`, `if`).
- It ensures that keywords are correctly tokenized as their specific types (e.g., `LET` for `let`) while treating non-keyword identifiers as `IDENT`.

---

#### **Structure**
Defined in `token.go`:
```go
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}
```
- **Key**: String representation of a keyword (e.g., `"let"`, `"fn"`).
- **Value**: Corresponding `TokenType` constant (e.g., `LET`, `FUNCTION`).

---

#### **How It Works**
1. **Lexer Tokenization**:
   - When the lexer encounters an identifier or keyword during tokenization, it uses the lookup table to decide the `TokenType`.

2. **Key Function**:
   ```go
   func LookupIdent(ident string) TokenType {
       if tok, ok := keywords[ident]; ok {
           return tok // If found in the map, it's a keyword
       }
       return IDENT // Otherwise, it's a user-defined identifier
   }
   ```

3. **Process**:
   - **Input**: A string literal (e.g., `"let"`, `"x"`).
   - **Output**:
     - If the string matches a keyword in `keywords`, it returns the corresponding `TokenType`.
     - Otherwise, it returns `IDENT` for user-defined identifiers.

---

#### **Example: "let x = 5;"**

1. **Processing `"let"`**:
   - The lexer reads `"let"` as a sequence of letters.
   - Calls `LookupIdent("let")`.
   - `"let"` exists in `keywords`, so it returns `TokenType` `LET`.
   - Emits the token:
     ```go
     {Type: LET, Literal: "let"}
     ```

2. **Processing `"x"`**:
   - The lexer reads `"x"` as a sequence of letters.
   - Calls `LookupIdent("x")`.
   - `"x"` is not in `keywords`, so it returns `TokenType` `IDENT`.
   - Emits the token:
     ```go
     {Type: IDENT, Literal: "x"}
     ```

3. **Processing Non-Keywords**:
   - Symbols like `"="` and `";"` are handled directly by `NextToken`, bypassing the lookup table.

---

#### **Why Use a Lookup Table?**
1. **Efficiency**:
   - The `map` data structure provides fast lookups for keywords.

2. **Scalability**:
   - Adding new keywords is straightforward—simply add entries to the `keywords` map.

3. **Clarity**:
   - Keeps logic clean by centralizing keyword classification.

---

Here’s a detailed breakdown of the **major components and interactions** in your lexer, focusing on elements you'll want to include in your UML diagram:

---

### **Major Components for the UML Diagram**

#### **1. `Lexer` (Defined in `lexer.go`)**
   - **Type**: Struct
   - **Fields**:
     - `input` (`string`): The source code being lexed.
     - `position` (`int`): Index of the current character in `input`.
     - `readPosition` (`int`): Index of the next character in `input`.
     - `ch` (`byte`): Current character being examined.
   - **Methods**:
     - `New(string) *Lexer`: Constructor to initialize a lexer.
     - `NextToken() token.Token`: Advances through the input and returns the next token.
     - `readChar()`: Moves to the next character in `input`.
     - `peekChar() byte`: Returns the next character without advancing `position`.
     - `skipWhitespace()`: Skips whitespace in the input.
     - `readIdentifier() string`: Reads a sequence of letters (e.g., variable names).
     - `readNumber() string`: Reads a sequence of digits (e.g., numbers).

---

#### **2. `Token` (Defined in `token.go`)**
   - **Type**: Struct
   - **Fields**:
     - `Type` (`TokenType`): Represents the type of the token (e.g., `LET`, `IDENT`, `INT`).
     - `Literal` (`string`): The value of the token (e.g., `"5"`, `"let"`).
   - **Enum**:
     - `TokenType`: Defined as constants for all valid token types (e.g., `LET`, `ASSIGN`, `PLUS`).
   - **Related Functions**:
     - `LookupIdent(string) TokenType`: Determines if a string is a keyword or an identifier.
     - `newToken(TokenType, byte) Token`: Helper function to create a token.

---

#### **3. `keywords` Map (Defined in `token.go`)**
   - **Purpose**: Differentiates between identifiers and reserved keywords.
   - **Type**: `map[string]TokenType`
   - **Interaction**:
     - Used in `LookupIdent` to determine the `TokenType` for strings.

---

#### **4. `REPL` (Defined in `repl.go`)**
   - **Type**: Functionality for interactive input/output.
   - **Methods**:
     - `Start(io.Reader, io.Writer)`: Reads input, tokenizes it using the lexer, and prints tokens.
   - **Interactions**:
     - Creates a `Lexer` instance.
     - Iterates through tokens using `NextToken()`.

---

#### **5. Workflow: Interaction Between Components**
For your UML diagram, show how these components interact:
1. **`REPL`**:
   - Accepts input from the user.
   - Passes input to `Lexer`.
2. **`Lexer`**:
   - Converts the input string into a sequence of tokens using `NextToken()`.
   - Uses helper methods (`readChar`, `readIdentifier`, `skipWhitespace`, etc.) to process input.
3. **`Token`**:
   - Represents each token generated by the lexer.
   - `TokenType` identifies the type of token (e.g., `LET`, `IDENT`).

---

### **UML Diagram Elements**
#### **Classes**
1. **`Lexer`**:
   - Attributes: `input`, `position`, `readPosition`, `ch`
   - Methods: `New`, `NextToken`, `readChar`, `peekChar`, `skipWhitespace`, `readIdentifier`, `readNumber`

2. **`Token`**:
   - Attributes: `Type`, `Literal`
   - Methods: `LookupIdent`, `newToken`

3. **`REPL`**:
   - Methods: `Start`

4. **`keywords` Map**:
   - Attributes: `map[string]TokenType`
   - Methods: None (but interacts with `LookupIdent`)

---

Would you like help drafting the UML diagram itself or a more specific visual representation of these interactions?