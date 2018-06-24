package symtab

type SymbolTable interface {
	Append(SymbolTable) (SymbolTable, error)
}

func New() SymbolTable {
	return nil
}

func NewVisitor() *SymbolTableVisitor {
	return &SymbolTableVisitor{}
}
