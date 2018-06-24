// Code generated from Concerto.G4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Concerto

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ConcertoListener is a complete listener for a parse tree produced by ConcertoParser.
type ConcertoListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterSimpleDeclaration is called when entering the simpleDeclaration production.
	EnterSimpleDeclaration(c *SimpleDeclarationContext)

	// EnterDeclarationConstructor is called when entering the declarationConstructor production.
	EnterDeclarationConstructor(c *DeclarationConstructorContext)

	// EnterArg is called when entering the arg production.
	EnterArg(c *ArgContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitSimpleDeclaration is called when exiting the simpleDeclaration production.
	ExitSimpleDeclaration(c *SimpleDeclarationContext)

	// ExitDeclarationConstructor is called when exiting the declarationConstructor production.
	ExitDeclarationConstructor(c *DeclarationConstructorContext)

	// ExitArg is called when exiting the arg production.
	ExitArg(c *ArgContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)
}
