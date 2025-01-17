// position and readPosition. Both will be used to access characters in input by
// using them as an index, e.g.: l.input[l.readPosition]. The reason for these two “pointers”
// pointing into our input string is the fact that we will need to be able to “peek” further into
// the input and look after the current character to see what comes up next. readPosition always
// points to the “next” character in the input. position points to the character in the input that
// corresponds to the ch byte.

// Lexer/lexer.go

package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // Tracks the location of the current character (ch) being processed.
	readPosition int  // Always points to the next character in the input string, allowing the lexer to "peek" ahead to make decisions (e.g., to distinguish between = and ==).
	ch           byte // current char under examination
}

// This is a constructor-like function in Go for initializing a new Lexer object.
// This defines a function named New that:
// 		Takes a single argument, input, of type string.
// 		Returns a pointer to a Lexer object (*Lexer).
func New(source_code_input string) *Lexer {
	l := &Lexer{input: source_code_input}
	l.readChar() // Fully initialize our *Lexer b4 anyone has the chance to call nextToken()
	return l
}

// We look at the current character under examination (l.ch) and return a token depending on which character it is.
// Before returning the token we advance our pointers into the input so when we call NextToken() again the l.ch field is already updated.
//
// LINE 61 -> "case '=':
//					tok = newToken(token.ASSIGN, l.ch)"
// This line ensures that the = character is correctly identified as an assignment operator and encapsulated as a token that the parser can understand.
// By encapsulating the type (ASSIGN) and literal value (=) in a token, the lexer translates raw input characters into structured information for the next stages of the interpreter/compiler.
// 		tokenType: Specifies the type of the token. Here, it is token.ASSIGN, which identifies the = token.
// 		ch: The character being processed, which is stored in l.ch (the current character the lexer is analyzing).
// The newly created token object is stored in tok, which is then returned by the NextToken function after advancing the lexer to the next character.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace() // Whitespace only acts as a separater of tokens in monkey language

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default: // We added a default branch to our switch statement, so we can check for identifiers whenever the l.ch is not one of the recognized characters.
		// Could probably generalize this by passing in the characteridentifying functions as arguments, but won’t, for simplicity’s sake and ease of understanding.
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()          // we use readIdentifier() to set the Literal field of our current token.
			tok.Type = token.LookupIdent(tok.Literal) // tell user-defined identifiers apart from language keywords NEED function that returns the correct TokenType for the token literal we have.
			return tok                                // The early exit here, our return tok statement, is necessary because when calling readIdentifier(), we call readChar() repeatedly and advance our readPosition and position fields past the last character of the current identifier.
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch) // we truly don’t know how to handle the current character and declare it as token.ILLEGAL.
		}
	}

	l.readChar()
	return tok
}

// Found in many parsers sometimes it’s called eatWhitespace and sometimes consumeWhitespace.
// Which characters these functions actually skip depends on the language being lexed.
// Some language implementations do create tokens for newline characters for example and throw parsing errors if they are not at the correct place in the stream of tokens.
// We skip over newline characters to make the parsing step later on a little easier.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// The purpose of readChar is to give us the next character and advance our position in the input string
// First checks if we have reached the end of input
// 		If so: sets .ch to - which is the ASCII code for the "NUL" character means:
//				1. we either havent read anything yet OR 2. End of the file for us
// 		else: sets l.ch to the next character by accessing l.input[l.readPosition]
func (l *Lexer) readChar() {
	// Check if at the end of the input
	if l.readPosition >= len(l.input) {
		l.ch = 0 // sets ch = 0 which is ASCII for 'Nul'
	} else {
		// Sets l.ch value == to the NEXT character in the input string
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// PeekChar() Similar to readChar except it wont increment l.position and l.readPosition
// PeekChar() is designed to peek ahead to the next character for certain TWO character tokens like '==' and '!='
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// SUPER SIMPLIFIED
// @ToDo Add in support for Floats, hex notation, octal notation etc...
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// Helper function just checks whether the given argument is a letter.
// Treat '_' as a letter allowing it in identifiers and keywords.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// Helper function just returns whether the passed in byte is a Latin digit between 0 and 9.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// Helper function that helps us with initializing these tokens.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
