package parser

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/dogslee/go-test-summery/node"
	"github.com/dogslee/go-test-summery/parser/visitor"
)

// WalfDir 深度遍输出全部目录
func WalfDirGroupFile(dir string, root node.Node) {

	testFileList := GetTestFileList(dir)

	if len(testFileList) > 0 {
		for _, testFile := range testFileList {
			// 创建一个TestFileNode
			testFileNode := node.NewTestFileNode(root, filepath.Base(testFile))

			// 判断函数中是否有测试和子测试
			vi := visitor.ParseFile(testFile)
			for _, fun := range vi.Func {
				funcName, commont := GetFuncNameAndComment(fun)
				// 判断是否为单元测试函数
				if strings.HasPrefix(funcName, "Test") {
					newTestNode := node.NewTestNode(testFileNode, funcName, commont)

					// 生成子测试函数SubTestNode list
					subTestList := GetSubTestNameAndComment(fun.Body, funcName, vi.Fset)
					for _, subTest := range subTestList {
						newSubTestNode := node.NewSubTestNode(newTestNode, funcName, subTest.TestName, subTest.FuncName, subTest.Comment)
						newTestNode.AddChild(newSubTestNode)
					}
					_, tn := node.SummeryNodeInfo(newTestNode)

					newTestNode.TestCnt += tn

					testFileNode.AddChild(newTestNode)
				}

			}

			root.AddChild(testFileNode)

		}
	}

	// 读取当前目录下的所有文件
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		if f.IsDir() && !strings.HasPrefix(f.Name(), ".") {
			nextDir := filepath.Join(dir, f.Name())
			// 创建一个DIRNode
			tmpDir := dir
			if strings.HasPrefix(tmpDir, "./") {
				tmpDir = tmpDir[2:]
				if len(tmpDir) > 2 && tmpDir[len(tmpDir)-1:] == "/" {
					tmpDir = tmpDir[:len(tmpDir)-1]
				}
			}
			dirName := strings.Replace(nextDir, tmpDir, "", -1)
			nextNode := node.NewDIRNode(root, dirName)

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
