// Chapter-1:
// - Going to write our own lexer.
// - Lexer will take source code as input and output the tokens that represent the source code.
// - Lexer reads in source code as a string, analyzes it character by character, and output the next
//   token it recognizes. 
// - ONLY need ONE method NextToken() which just outputs the next token
// WE SHOULD ADD FILE NAMES AND LINE NUMBERS TO TOKENS BUT HERE IT DOESNT MAKE SENSE TO DO SO.

// token/token.go // Go back to page 17 and add file names and line numbers
package token

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
}

// LookupIdent checks the keywords table to see whether the given identifier is in fact a keyword.
// If it is, it returns the keyword’s TokenType constant.
// If it isn’t, we just get back token.IDENT, which is the TokenType for all user-defined identifiers.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
	return tok
	}
	return IDENT
}

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

// We have a LIMITED number of token types so we can define them as CONSTANTS
const (
	ILLEGAL = "ILLEGAL" // Illegal signifies a character or token we dont know about
	EOF = "EOF" // End of file

	// Identifiers + Literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT = "INT" // 13456

	// Operators
	ASSIGN = "="
	PLUS = "+"

	// Delimeters 
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)