package node

import "fmt"

// SubTestNode is a sub test node
type SubTestNode struct {
	// 父节点的名称
	ParentName string
	TestName   string
	FuncName   string
	Comment    string
}

func (s *SubTestNode) GetChildren() []Node {
	return nil
}

func (s *SubTestNode) GetName() string {
	ret := ""
	if len(s.Comment) > 0 {
		ret = fmt.Sprintf("%s:%s/%s", s.Comment, s.ParentName, s.TestName)
	} else {
		ret = fmt.Sprintf("%s/%s", s.ParentName, s.TestName)
	}
	return ret
}

func (s *SubTestNode) AddChild(node Node) {
	panic("implement me")
}

func (s *SubTestNode) IsLeaf() bool {
	return true
}

// NewSubTestNode is a constructor of SubTestNode
func NewSubTestNode(parentName, testName, funcName, comment string) *SubTestNode {
	return &SubTestNode{
		ParentName: parentName,
		TestName:   testName,
		FuncName:   funcName,
		Comment:    comment,
	}
}
