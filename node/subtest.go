package node

// SubTestNode is a sub test node
type SubTestNode struct {
	Parent Node
	// 父节点的名称
	ParentName string
	TestName   string // 测试名称
	FuncName   string // 函数名称 默认测试名称和函数名称是一样的
	Comment    string
}

func (s *SubTestNode) GetParent() Node {
	return s.Parent
}

func (s *SubTestNode) GetChildren() []Node {
	return nil
}

func (s *SubTestNode) GetName() string {
	return genTestInfo(s.ParentName, s.FuncName, s.Comment)
}

func (s *SubTestNode) AddChild(node Node) {
	panic("implement me")
}

func (s *SubTestNode) IsLeaf() bool {
	return true
}

// NewSubTestNode is a constructor of SubTestNode
func NewSubTestNode(p Node, parentName, testName, funcName, comment string) *SubTestNode {
	return &SubTestNode{
		Parent:     p,
		ParentName: parentName,
		TestName:   testName,
		FuncName:   funcName,
		Comment:    comment,
	}
}
