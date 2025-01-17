// ast/ast.go

// The AST we are going to construct consists solely of Nodes that are connected to each other - it’s a tree after all.
// Some of these nodes implement the Statement and some the Expression interface.
//
// These interfaces only contain dummy methods called statementNode and expressionNode respectively.
// They are not strictly necessary but help us by guiding the Go compiler and possibly causing it
// to throw errors when we use a Statement where an Expression should’ve been used, and vice versa.

package ast

import (
	// "bytes"
	"monkey/token"
	// "strings"
)

// Every node in our AST has to implement the Node interface, meaning it has to provide a TokenLiteral() method that returns
// the literal value of the token it’s associated with. TokenLiteral() will be used only for debugging and testing.
// The AST we are going to construct
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode() // dummy method
}

type Expression interface {
	Node
	expressionNode() // dummy method
}

// This Program node is going to be the root node of every AST our parser produces.
// Every valid Monkey program is a series of statements.
// These statements are contained in the Program.Statements,
// which is just a slice of AST nodes that implement the Statement interface.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier // Name to hold the identifier of the binding
	Value Expression  // Value for the expression that produces the value
}

// The two methods statementNode and TokenLiteral satisfy the Statement and Node interfaces respectively.
func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// To keep the number of different node types small, we’ll use Identifier here to represent the name
// in a variable binding and later reuse it, to represent an identifier as part of or as a complete expression.
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (ls *Identifier) expressionNode()      {}
func (ls *Identifier) TokenLiteral() string { return i.Token.Literal }
