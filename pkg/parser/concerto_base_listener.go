// Code generated from Concerto.G4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Concerto

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseConcertoListener is a complete listener for a parse tree produced by ConcertoParser.
type BaseConcertoListener struct{}

var _ ConcertoListener = &BaseConcertoListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseConcertoListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseConcertoListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseConcertoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseConcertoListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseConcertoListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseConcertoListener) ExitProg(ctx *ProgContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseConcertoListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseConcertoListener) ExitStatement(ctx *StatementContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseConcertoListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseConcertoListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterSimpleDeclaration is called when production simpleDeclaration is entered.
func (s *BaseConcertoListener) EnterSimpleDeclaration(ctx *SimpleDeclarationContext) {}

// ExitSimpleDeclaration is called when production simpleDeclaration is exited.
func (s *BaseConcertoListener) ExitSimpleDeclaration(ctx *SimpleDeclarationContext) {}

// EnterDeclarationConstructor is called when production declarationConstructor is entered.
func (s *BaseConcertoListener) EnterDeclarationConstructor(ctx *DeclarationConstructorContext) {}

// ExitDeclarationConstructor is called when production declarationConstructor is exited.
func (s *BaseConcertoListener) ExitDeclarationConstructor(ctx *DeclarationConstructorContext) {}

// EnterArg is called when production arg is entered.
func (s *BaseConcertoListener) EnterArg(ctx *ArgContext) {}

// ExitArg is called when production arg is exited.
func (s *BaseConcertoListener) ExitArg(ctx *ArgContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseConcertoListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseConcertoListener) ExitExpression(ctx *ExpressionContext) {}
