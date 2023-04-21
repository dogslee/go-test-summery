package node

type DirNode struct {
	DirName  string
	Children []Node
}

func (d *DirNode) GetChildren() []Node {
	ret := make([]Node, 0)
	for _, child := range d.Children {
		ret = append(ret, child)
	}
	return ret
}

func (d *DirNode) GetName() string {
	return d.DirName
}

func (d *DirNode) AddChild(node Node) {
	d.Children = append(d.Children, node)
}

func (d *DirNode) IsLeaf() bool {
	return false
}

// NewDIRNode is a constructor of DIRNode
func NewDIRNode(dirName string) *DirNode {
	return &DirNode{
		DirName: dirName,
	}
}
