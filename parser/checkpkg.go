package parser

import (
	"go/parser"
	"go/token"
	"io/fs"
)

// CheckPkg 检查包中是否含有内容
func CheckPkg(dir string) bool {
	pkgMap, err := parser.ParseDir(token.NewFileSet(), dir, func(_ fs.FileInfo) bool {
		return true
	}, 4)
	return len(pkgMap) > 0 && err == nil
}
