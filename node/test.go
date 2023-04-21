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
	ret := ""
	if len(t.Comment) > 0 {
		ret = fmt.Sprintf("%s:%s", t.Comment, t.FuncName)
	} else {
		ret = t.FuncName
	}
	return ret
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
