package ast

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	// statementNode is a marker method. It can be used to differentiate between
	// different types of statements.
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}
