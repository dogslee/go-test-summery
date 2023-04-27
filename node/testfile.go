package node

type TestFileNode struct {
	FileName    string
	Parent      Node
	Children    []Node
	TestCnt     int
	TestPassCnt int
	TestFailCnt int
}

func (t *TestFileNode) GetParent() Node {
	return t.Parent
}

func (t *TestFileNode) GetChildren() []Node {
	return t.Children
}

func (t *TestFileNode) GetName() string {
	return t.FileName
}

func (t *TestFileNode) AddChild(node Node) {
	t.Children = append(t.Children, node)
}

func (t *TestFileNode) IsLeaf() bool {
	return len(t.Children) == 0
}

// NewTestFileNode is a constructor of TestFileNode
func NewTestFileNode(p Node, fileName string) *TestFileNode {
	return &TestFileNode{
		Parent:   p,
		FileName: fileName,
	}
}
