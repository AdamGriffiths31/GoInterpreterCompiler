package compiler

type SymbolScope string

const (
	GlobalScope SymbolScope = "GLOBAL"
)

type Symbol struct {
	Name  string
	Scope SymbolScope
	Index int
}

type SymbolTable struct {
	store          map[string]Symbol
	numDefinitions int
}

func NewSymbolTable() *SymbolTable {
	s := make(map[string]Symbol)
	return &SymbolTable{store: s}
}

func (t *SymbolTable) Define(name string) Symbol {
	symbol := Symbol{Name: name, Index: t.numDefinitions, Scope: GlobalScope}
	t.store[name] = symbol
	t.numDefinitions++
	return symbol
}

func (t *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := t.store[name]
	return obj, ok
}
