package visitor

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
)

type Visitor struct {
	Fset   *token.FileSet
	Func   []*ast.FuncDecl
	Struct []*ast.StructType
	Ids    []*ast.Ident
	Type   []*ast.TypeSpec
}

func (v *Visitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.FuncDecl:
		v.Func = append(v.Func, n)
	case *ast.StructType:
		v.Struct = append(v.Struct, n)
	case *ast.Ident:
		v.Ids = append(v.Ids, n)
	case *ast.TypeSpec:
		v.Type = append(v.Type, n)
	}
	return v
}

func ParseDir(fileDir string) *Visitor {
	v := &Visitor{}

	// positions are relative to fileSet
	fileSet := token.NewFileSet()
	v.Fset = fileSet

	pkgMap, err := parser.ParseDir(v.Fset, fileDir, func(_ fs.FileInfo) bool {
		return true
	}, 4)
	if err != nil {
		panic(err)
	}

	for _, pkg := range pkgMap {
		// dfs get all node
		for _, f := range pkg.Files {
			ast.Walk(v, f)
		}
	}
	return v
}

func ParseFile(file string) *Visitor {
	v := &Visitor{}

	// positions are relative to fileSet
	fileSet := token.NewFileSet()
	v.Fset = fileSet

	f, err := parser.ParseFile(v.Fset, file, nil, 4)
	if err != nil {
		panic(err)
	}

	ast.Walk(v, f)
	return v
}
