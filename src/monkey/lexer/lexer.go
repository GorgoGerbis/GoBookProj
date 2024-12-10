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
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // Fully initialize our *Lexer b4 anyone has teh chance to call nextToken()
	return l
}

// The purpose of readChar is to give us the next character and advance our position in the input string
// First checks if we have reached the end of input
// 		If so: sets .ch to - which is the ASCII code for the "NUL" character means:
//				1. we either havent read anything yet OR 2. End of the file for us
// 		else: sets l.ch to the next character by accessing l.input[l.readPosition]
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// We look at the current character under examination (l.ch) and return a token depending on which character it is.
// Before returning the token we advance our pointers into the input so when we call NextToken() again the l.ch field is already updated.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

// Helper function that helps us with initializing these tokens.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
