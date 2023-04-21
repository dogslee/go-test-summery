package parser

import (
	"go-test-summery/node"
	"go-test-summery/parser/visitor"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

// WalfDir 深度遍输出全部目录
func WalfDir(dir string, root node.Node) {
	if dir != "log" {
		// 判断当前目录下是否包含_test.go文件, 如果包含则创建一个TestNode
		if CheckTestFile(dir) {
			vi := visitor.ParseDir(dir)
			for _, fun := range vi.Func {
				funcName, commont := GetFuncNameAndComment(fun)
				// 判断是否为单元测试函数
				if strings.HasPrefix(funcName, "Test") {
					newTestNode := node.NewTestNode(funcName, commont)
					root.AddChild(newTestNode)

					// 生成子测试函数SubTestNode list
					subTestList := GetSubTestNameAndComment(fun.Body, funcName, vi.Fset)
					for _, subTest := range subTestList {
						newSubTestNode := node.NewSubTestNode(funcName, subTest.TestName, subTest.FuncName, subTest.Comment)
						newTestNode.AddChild(newSubTestNode)
					}
				}
			}
		}
	}
	// 读取当前目录下的所有文件
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		if f.IsDir() && f.Name() != "log" {
			nextDir := filepath.Join(dir, f.Name())

			// 创建一个DIRNode
			dirName := strings.Replace(nextDir, dir, "", -1)
			nextNode := node.NewDIRNode(dirName)

			root.AddChild(nextNode)
			WalfDir(nextDir, nextNode)
		}
	}
}

// WalfDir 深度遍输出全部目录
func WalfDirGroupFile(dir string, root node.Node) {

	testFileList := GetTestFileList(dir)

	if len(testFileList) > 0 {
		for _, testFile := range testFileList {
			// 创建一个TestFileNode
			testFileNode := node.NewTestFileNode(filepath.Base(testFile))
			root.AddChild(testFileNode)

			// 判断函数中是否有测试和子测试
			vi := visitor.ParseFile(testFile)
			for _, fun := range vi.Func {
				funcName, commont := GetFuncNameAndComment(fun)
				// 判断是否为单元测试函数
				if strings.HasPrefix(funcName, "Test") {
					newTestNode := node.NewTestNode(funcName, commont)
					testFileNode.AddChild(newTestNode)

					// 生成子测试函数SubTestNode list
					subTestList := GetSubTestNameAndComment(fun.Body, funcName, vi.Fset)
					for _, subTest := range subTestList {
						newSubTestNode := node.NewSubTestNode(funcName, subTest.TestName, subTest.FuncName, subTest.Comment)
						newTestNode.AddChild(newSubTestNode)
					}
				}
			}
		}
	}

	// 读取当前目录下的所有文件
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		if f.IsDir() && f.Name() != "log" {
			nextDir := filepath.Join(dir, f.Name())

			// 创建一个DIRNode
			dirName := strings.Replace(nextDir, dir, "", -1)
			nextNode := node.NewDIRNode(dirName)

			root.AddChild(nextNode)
			WalfDirGroupFile(nextDir, nextNode)
		}
	}
}

// 判断当前目录下是否包含_test.go文件
func CheckTestFile(dir string) bool {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.Name() == "log" {
			continue
		}
		if strings.HasSuffix(f.Name(), "_test.go") {
			return true
		}
	}
	return false
}

func GetTestFileList(dir string) []string {
	fileList := make([]string, 0)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.Name() == "log" {
			continue
		}
		if strings.HasSuffix(f.Name(), "_test.go") {
			fileList = append(fileList, filepath.Join(dir, f.Name()))
		}
	}
	return fileList
}
