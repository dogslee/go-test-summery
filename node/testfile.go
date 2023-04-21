package node

type TestFileNode struct {
	FileName string
	Children []Node
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
func NewTestFileNode(fileName string) *TestFileNode {
	return &TestFileNode{
		FileName: fileName,
	}
}
