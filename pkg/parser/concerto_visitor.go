// Code generated from Concerto.G4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Concerto

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ConcertoParser.
type ConcertoVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ConcertoParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by ConcertoParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by ConcertoParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by ConcertoParser#simpleDeclaration.
	VisitSimpleDeclaration(ctx *SimpleDeclarationContext) interface{}

	// Visit a parse tree produced by ConcertoParser#declarationConstructor.
	VisitDeclarationConstructor(ctx *DeclarationConstructorContext) interface{}

	// Visit a parse tree produced by ConcertoParser#arg.
	VisitArg(ctx *ArgContext) interface{}

	// Visit a parse tree produced by ConcertoParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}
}
