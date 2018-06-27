package symtab

type SymbolTable interface {
	Append(SymbolTable) (SymbolTable, error)
	Error() error
}

type BaseSymbolTable struct {
	Err error
}

func (s BaseSymbolTable) Append(tab SymbolTable) (SymbolTable, error) {
	return s, nil
}

func (s BaseSymbolTable) Error() error {
	return s.Err
}

func New() SymbolTable {
	return nil
}

func NewVisitor() *SymbolTableVisitor {
	return &SymbolTableVisitor{}
}
