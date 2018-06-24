// Code generated from Concerto.G4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Concerto

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseConcertoVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseConcertoVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConcertoVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConcertoVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConcertoVisitor) VisitSimpleDeclaration(ctx *SimpleDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConcertoVisitor) VisitDeclarationConstructor(ctx *DeclarationConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConcertoVisitor) VisitArg(ctx *ArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConcertoVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}
