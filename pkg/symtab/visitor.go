package symtab

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/koki/concerto/pkg/parser"
)

type SymbolTableVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *SymbolTableVisitor) VisitProg(ctx *parser.ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SymbolTableVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SymbolTableVisitor) VisitDeclaration(ctx *parser.DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SymbolTableVisitor) VisitSimpleDeclaration(ctx *parser.SimpleDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SymbolTableVisitor) VisitDeclarationConstructor(ctx *parser.DeclarationConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SymbolTableVisitor) VisitArg(ctx *parser.ArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SymbolTableVisitor) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}
