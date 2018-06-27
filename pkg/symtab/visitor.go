package symtab

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/golang/glog"
	"github.com/koki/concerto/pkg/parser"
)

var _ parser.ConcertoVisitor = &SymbolTableVisitor{}

const SYMBOL_TABLE_VISITOR = "SYMBOL_TABLE_VISITOR"

type SymbolTableVisitor struct {
	antlr.ParseTreeVisitor
}

func (v *SymbolTableVisitor) VisitProg(ctx *parser.ProgContext) interface{} {
	glog.V(3).Infof("[%s] entering program context", SYMBOL_TABLE_VISITOR)

	v.VisitPackageClause(ctx.PackageClause().(*parser.PackageClauseContext))
	v.VisitImportDecl(ctx.ImportDecl().(*parser.ImportDeclContext))

	for _, decl := range ctx.AllTopLevelDecl() {
		v.VisitTopLevelDecl(decl.(*parser.TopLevelDeclContext))
	}

	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitPackageClause(ctx *parser.PackageClauseContext) interface{} {
	glog.V(3).Infof("[%s] entering package clause context [%s]", SYMBOL_TABLE_VISITOR, ctx.IDENTIFIER())
	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitImportDecl(ctx *parser.ImportDeclContext) interface{} {
	glog.V(3).Infof("[%s] entering import decl context", SYMBOL_TABLE_VISITOR)

	for _, spec := range ctx.AllImportSpec() {
		v.VisitImportSpec(spec.(*parser.ImportSpecContext))
	}

	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitImportSpec(ctx *parser.ImportSpecContext) interface{} {
	valx := ""
	val := ctx.STRING_LIT()
	if val != nil {
		valx = val.GetText()
	}
	glog.V(3).Infof("[%s] entering import spec context [%s] %s", SYMBOL_TABLE_VISITOR, ctx.IDENTIFIER().GetText(), valx)
	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitEos(ctx *parser.EosContext) interface{} {
	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitTopLevelDecl(ctx *parser.TopLevelDeclContext) interface{} {
	glog.V(3).Infof("[%s] entering top level decl context", SYMBOL_TABLE_VISITOR)

	if fn := ctx.FuncDecl(); fn != nil {
		v.VisitFuncDecl(fn.(*parser.FuncDeclContext))
	}

	if m := ctx.MethodDecl(); m != nil {
		v.VisitMethodDecl(m.(*parser.MethodDeclContext))
	}

	if d := ctx.Declaration(); d != nil {
		v.VisitDeclaration(d.(*parser.DeclarationContext))
	}

	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitFuncDecl(ctx *parser.FuncDeclContext) interface{} {
	glog.V(3).Infof("[%s] entering func decl context", SYMBOL_TABLE_VISITOR)

	v.VisitMethodSpec(ctx.MethodSpec().(*parser.MethodSpecContext))
	v.VisitBlock(ctx.Block().(*parser.BlockContext))

	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitMethodDecl(ctx *parser.MethodDeclContext) interface{} {
	glog.V(3).Infof("[%s] entering method decl context for type %s", SYMBOL_TABLE_VISITOR, ctx.IDENTIFIER())

	v.VisitMethodSpec(ctx.MethodSpec().(*parser.MethodSpecContext))
	v.VisitBlock(ctx.Block().(*parser.BlockContext))

	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	glog.V(3).Infof("[%s] entering block context", SYMBOL_TABLE_VISITOR)

	for _, stat := range ctx.AllStatement() {
		v.VisitStatement(stat.(*parser.StatementContext))
	}
	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	glog.V(3).Infof("[%s] entering statement", SYMBOL_TABLE_VISITOR)

	if sd := ctx.StatementDecl(); sd != nil {
		v.VisitStatementDecl(sd.(*parser.StatementDeclContext))
	}

	if re := ctx.ReturnExpr(); re != nil {
		v.VisitReturnExpr(re.(*parser.ReturnExprContext))
	}

	if expr := ctx.Expression(); expr != nil {
		v.VisitExpression(expr.(*parser.ExpressionContext))
	}

	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitReturnExpr(ctx *parser.ReturnExprContext) interface{} {
	glog.V(3).Infof("[%s] entering return expr context", SYMBOL_TABLE_VISITOR)
	v.VisitExpression(ctx.Expression().(*parser.ExpressionContext))
	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitStatementDecl(ctx *parser.StatementDeclContext) interface{} {
	glog.V(3).Infof("[%s] entering statement decl context", SYMBOL_TABLE_VISITOR)
	v.VisitVarDecl(ctx.VarDecl().(*parser.VarDeclContext))
	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	glog.V(3).Infof("[%s] entering expression context", SYMBOL_TABLE_VISITOR)
	if pe := ctx.PrimaryExpr(); pe != nil {
		v.VisitPrimaryExpr(pe.(*parser.PrimaryExprContext))
	}

	if ctx.OR() != nil {
		v.VisitOP(parser.ConcertoParserOR, ctx.AllExpression())
	}

	if ctx.AND() != nil {
		v.VisitOP(parser.ConcertoParserAND, ctx.AllExpression())
	}

	if ctx.LT() != nil {
		v.VisitOP(parser.ConcertoParserLT, ctx.AllExpression())
	}

	if ctx.LE() != nil {
		v.VisitOP(parser.ConcertoParserLE, ctx.AllExpression())
	}

	if ctx.GT() != nil {
		v.VisitOP(parser.ConcertoParserGT, ctx.AllExpression())
	}

	if ctx.GE() != nil {
		v.VisitOP(parser.ConcertoParserGE, ctx.AllExpression())
	}

	if ctx.PLUS() != nil {
		v.VisitOP(parser.ConcertoParserPLUS, ctx.AllExpression())
	}

	if ctx.MINUS() != nil {
		v.VisitOP(parser.ConcertoParserMINUS, ctx.AllExpression())
	}

	if ctx.BITWISE_OR() != nil {
		v.VisitOP(parser.ConcertoParserBITWISE_OR, ctx.AllExpression())
	}

	if ctx.CARET() != nil {
		v.VisitOP(parser.ConcertoParserCARET, ctx.AllExpression())
	}

	if ctx.MUL() != nil {
		v.VisitOP(parser.ConcertoParserMUL, ctx.AllExpression())
	}

	if ctx.DIV() != nil {
		v.VisitOP(parser.ConcertoParserDIV, ctx.AllExpression())
	}

	if ctx.REM() != nil {
		v.VisitOP(parser.ConcertoParserREM, ctx.AllExpression())
	}

	if ctx.LSHIFT() != nil {
		v.VisitOP(parser.ConcertoParserLSHIFT, ctx.AllExpression())
	}

	if ctx.RSHIFT() != nil {
		v.VisitOP(parser.ConcertoParserRSHIFT, ctx.AllExpression())
	}

	if ctx.BITWISE_OR() != nil {
		v.VisitOP(parser.ConcertoParserBITWISE_AND, ctx.AllExpression())
	}

	if ctx.AND_NOT() != nil {
		v.VisitOP(parser.ConcertoParserAND_NOT, ctx.AllExpression())
	}

	if ctx.ASSIGN() != nil {
		v.VisitOP(parser.ConcertoParserASSIGN, ctx.AllExpression())
	}

	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitOP(op int, exprs []parser.IExpressionContext) interface{} {
	glog.V(3).Infof("[%s] entering OP context", SYMBOL_TABLE_VISITOR)
	for _, expr := range exprs {
		v.VisitExpression(expr.(*parser.ExpressionContext))
	}
	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitPrimaryExpr(ctx *parser.PrimaryExprContext) interface{} {
	glog.V(3).Infof("[%s] entering primary expr context", SYMBOL_TABLE_VISITOR)
	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitOperand(ctx *parser.OperandContext) interface{} {
	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitBasicLit(ctx *parser.BasicLitContext) interface{} {
	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitOperandName(ctx *parser.OperandNameContext) interface{} {
	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitQualifiedIdent(ctx *parser.QualifiedIdentContext) interface{} {
	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitVarDecl(ctx *parser.VarDeclContext) interface{} {
	glog.V(3).Infof("[%s] entering var decl context", SYMBOL_TABLE_VISITOR)

	name := ""
	typ := ""

	ids := ctx.AllIDENTIFIER()
	name = ids[0].GetText()
	if len(ids) == 2 {
		typ = ids[1].GetText()
	} else {
		v.VisitFuncCallSpec(ctx.FuncCallSpec().(*parser.FuncCallSpecContext))
	}

	glog.V(3).Infof("[%s] var %s %s", SYMBOL_TABLE_VISITOR, name, typ)

	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitDeclaration(ctx *parser.DeclarationContext) interface{} {
	glog.V(3).Infof("[%s] entering declaration context", SYMBOL_TABLE_VISITOR)
	v.VisitTypeDecl(ctx.TypeDecl().(*parser.TypeDeclContext))
	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitTypeDecl(ctx *parser.TypeDeclContext) interface{} {
	glog.V(3).Infof("[%s] entering type decl context", SYMBOL_TABLE_VISITOR)

	if sd := ctx.StructDecl(); sd != nil {
		v.VisitStructDecl(sd.(*parser.StructDeclContext))
	}

	if intDecl := ctx.InterfaceDecl(); intDecl != nil {
		v.VisitInterfaceDecl(intDecl.(*parser.InterfaceDeclContext))
	}

	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitStructDecl(ctx *parser.StructDeclContext) interface{} {
	glog.V(3).Infof("[%s] entering struct decl context", SYMBOL_TABLE_VISITOR)
	name := ""
	implements := []string{}

	for i, field := range ctx.AllIDENTIFIER() {
		if i == 0 {
			name = field.GetText()
			continue
		}
		implements = append(implements, field.GetText())
	}

	for _, tr := range ctx.AllTypeSpec() {
		v.VisitTypeSpec(tr.(*parser.TypeSpecContext))
	}

	glog.V(3).Infof("[%s] struct %s implements %+v", SYMBOL_TABLE_VISITOR, name, implements)

	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitTypeSpec(ctx *parser.TypeSpecContext) interface{} {
	glog.V(3).Infof("[%s] entering type spec context %s", SYMBOL_TABLE_VISITOR, ctx.IDENTIFIER())
	v.VisitTypeRule(ctx.TypeRule().(*parser.TypeRuleContext))
	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitInterfaceDecl(ctx *parser.InterfaceDeclContext) interface{} {
	glog.V(3).Infof("[%s] entering interface decl context %s", SYMBOL_TABLE_VISITOR, ctx.IDENTIFIER())
	for _, m := range ctx.AllMethodSpec() {
		v.VisitMethodSpec(m.(*parser.MethodSpecContext))
	}
	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitMethodSpec(ctx *parser.MethodSpecContext) interface{} {
	glog.V(3).Infof("[%s] entering method spec context", SYMBOL_TABLE_VISITOR)
	v.VisitFuncSpec(ctx.FuncSpec().(*parser.FuncSpecContext))
	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitFuncCallSpec(ctx *parser.FuncCallSpecContext) interface{} {
	glog.V(3).Infof("[%s] entering func call spec context of %s", SYMBOL_TABLE_VISITOR, ctx.IDENTIFIER())
	for _, arg := range ctx.AllFuncCallArg() {
		v.VisitFuncCallArg(arg.(*parser.FuncCallArgContext))
	}
	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitFuncCallArg(ctx *parser.FuncCallArgContext) interface{} {
	glog.V(3).Infof("[%s] entering func call arg context", SYMBOL_TABLE_VISITOR)
	v.VisitExpression(ctx.Expression().(*parser.ExpressionContext))
	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitFuncSpec(ctx *parser.FuncSpecContext) interface{} {
	glog.V(3).Infof("[%s] entering func spec context %s", SYMBOL_TABLE_VISITOR, ctx.IDENTIFIER())
	for _, arg := range ctx.AllFuncArg() {
		v.VisitFuncArg(arg.(*parser.FuncArgContext))
	}
	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitFuncArg(ctx *parser.FuncArgContext) interface{} {
	glog.V(3).Infof("[%s] entering func arg context", SYMBOL_TABLE_VISITOR)
	if id := ctx.IDENTIFIER(); id != nil {
		glog.V(3).Infof("[%s] func arg id %s", SYMBOL_TABLE_VISITOR, id.GetText())
	}
	v.VisitTypeRule(ctx.TypeRule().(*parser.TypeRuleContext))
	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitTypeRule(ctx *parser.TypeRuleContext) interface{} {
	glog.V(3).Infof("[%s] entering type rule context", SYMBOL_TABLE_VISITOR)
	first := ""
	second := ""
	for _, idNode := range ctx.AllIDENTIFIER() {
		if first == "" {
			first = idNode.GetText()
		} else {
			second = idNode.GetText()
		}
	}
	if first != "" || second != "" {
		glog.V(3).Infof("[%s] argtype %s.%s", SYMBOL_TABLE_VISITOR, first, second)
	}
	if ar := ctx.ArrayType(); ar != nil {
		v.VisitArrayType(ar.(*parser.ArrayTypeContext))
	}
	if mp := ctx.MapType(); mp != nil {
		v.VisitMapType(mp.(*parser.MapTypeContext))
	}
	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitArrayType(ctx *parser.ArrayTypeContext) interface{} {
	glog.V(3).Infof("[%s] entering array type context", SYMBOL_TABLE_VISITOR)
	ln := ""
	ids := ctx.AllIDENTIFIER()
	at := ""
	if len(ids) == 2 {
		ln = "[" + ids[0].GetText() + "]"
		at = ids[1].GetText()
	} else {
		ln = "[]" //[]arr
		at = ids[0].GetText()
	}
	if ctx.Star() != nil {
		ln = "[*]" //[*]arr
	}
	if i := ctx.INT_LIT(); i != nil {
		ln = "[" + i.GetText() + "]"
	}

	glog.V(3).Infof("[%s] array type %s of len %s", SYMBOL_TABLE_VISITOR, at, ln)

	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitMapType(ctx *parser.MapTypeContext) interface{} {
	glog.V(3).Infof("[%s] entering map type context", SYMBOL_TABLE_VISITOR)
	for _, tr := range ctx.AllTypeRule() {
		v.VisitTypeRule(tr.(*parser.TypeRuleContext))
	}
	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitStar(ctx *parser.StarContext) interface{} {
	return &BaseSymbolTable{nil}
}

func (v *SymbolTableVisitor) VisitArraySelection(ctx *parser.ArraySelectionContext) interface{} {
	return &BaseSymbolTable{nil}
}
