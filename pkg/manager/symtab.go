package manager

import (
	"github.com/koki/concerto/pkg/symtab"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func ExtractSymbolTable(trees []antlr.ParserRuleContext) (symtab.SymbolTable, error) {
	symTab := symtab.New()
	symTabVisitor := symtab.NewVisitor()

	// generate symbol table
	for _, tree := range trees {
		tab := tree.Accept(symTabVisitor)
		//TODO: Error context for visitors
		//if err != nil {
		//	return nil, err
		//}

		var err error
		symTab, err = symTab.Append(tab.(symtab.SymbolTable))
		if err != nil {
			return nil, err
		}
	}
	return symTab, nil
}
