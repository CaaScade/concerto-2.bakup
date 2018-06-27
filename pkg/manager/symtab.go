package manager

import (
	"github.com/koki/concerto/pkg/symtab"
	"github.com/koki/concerto/pkg/util"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func ExtractSymbolTable(trees []antlr.ParserRuleContext) (symtab.SymbolTable, error) {
	symTab := symtab.New()
	symTabVisitor := symtab.NewVisitor()

	// generate symbol table
	for _, tree := range trees {
		visitorContext := tree.Accept(symTabVisitor)
		if err := visitorContext.(util.VisitorContext).Error(); err != nil {
			return nil, err
		}

		//tab := visitorContext.(symtab.SymbolTable)

		//var err error
		//symTab, err = symTab.Append(tab.(symtab.SymbolTable))
		//if err != nil {
		//	return nil, err
		//}
	}
	return symTab, nil
}
