package parser

import (
	"go/ast"
	"go/token"
	"strings"

	"github.com/dogslee/go-test-summery/node"
)

// GetFuncNameAndComment 获取单元测试的方法名和和注释
func GetFuncNameAndComment(v *ast.FuncDecl) (funcName, comment string) {
	if v.Doc != nil {
		comment = v.Doc.Text()
		comment = strings.Replace(comment, "\n", "", -1)
	}
	funcName = v.Name.Name
	return funcName, comment
}

// GetSubTestNameAndComment 获取子测试列表
func GetSubTestNameAndComment(body *ast.BlockStmt, parentName string, fset *token.FileSet) []*node.SubTestNode {
	ret := make([]*node.SubTestNode, 0)
	for _, stmt := range body.List {
		if stmt, ok := stmt.(*ast.ExprStmt); ok {
			if call, ok := stmt.X.(*ast.CallExpr); ok {
				if fun, ok := call.Fun.(*ast.SelectorExpr); ok {
					if fun.Sel.Name == "Run" {
						if len(call.Args) == 2 {
							if arg0, ok := call.Args[0].(*ast.BasicLit); ok {
								if arg1, ok := call.Args[1].(*ast.Ident); ok {

									subTestName := strings.Replace(arg0.Value, "\"", "", -1)
									subFuncName := arg1.String()
									subTestComment := ""
									if arg1.Obj != nil && arg1.Obj.Decl != nil {
										declObj := arg1.Obj.Decl
										if decl, ok := declObj.(*ast.FuncDecl); ok {
											subTestComment = decl.Doc.Text()
										}

									}
									subTestComment = strings.Replace(subTestComment, "\n", "", -1)
									ret = append(ret, &node.SubTestNode{
										ParentName: parentName,
										TestName:   subTestName,
										FuncName:   subFuncName,
										Comment:    subTestComment,
									})
									// fmt.Println("subTestName", subTestName)
									// fmt.Println("subFuncName", subFuncName)
									// fmt.Println("subTestComment", subTestComment)
									//ast.Print(fset, stmt)
								}
							}
						}
					}
				}
			}
		}
	}
	return ret
}
