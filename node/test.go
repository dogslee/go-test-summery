package node

import "fmt"

// TestNode is a test node
type TestNode struct {
	FuncName string
	Comment  string
	Children []Node
}

func (t *TestNode) GetChildren() []Node {
	return t.Children
}

func (t *TestNode) GetName() string {
	return genTestInfo("", t.FuncName, t.Comment)
}

func (t *TestNode) AddChild(node Node) {
	t.Children = append(t.Children, node)
}

func (t *TestNode) IsLeaf() bool {
	return len(t.Children) == 0
}

// NewTestNode is a constructor of TestNode
func NewTestNode(funcName, comment string) Node {
	return &TestNode{
		FuncName: funcName,
		Comment:  comment,
	}
}

// genTestInfo 生成测试信息
func genTestInfo(testName, funcName, comment string) string {
	ret := ""
	if len(comment) > 0 {
		ret = fmt.Sprintf("%s:%s", funcName, comment)
	} else {
		ret = funcName
	}

	if testName != "" {
		ret = fmt.Sprintf("%s/%s", testName, ret)
	}
	return ret
}
