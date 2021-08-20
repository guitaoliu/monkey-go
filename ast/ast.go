package ast

import "github.com/guitaoliu/monkey-go/token"

type Node interface {
	TokenLiterial() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Root of the program
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiterial() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiterial()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()        {}
func (ls *LetStatement) TokenLiterial() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()       {}
func (i *Identifier) TokenLiterial() string { return i.Token.Literal }
