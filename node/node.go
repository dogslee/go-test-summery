package node

type Node interface {
	// 获取父节点
	GetParent() Node

	// 获取子节点list
	GetChildren() []Node

	// 获取节点名称
	GetName() string

	// 添加子节点
	AddChild(node Node)

	// 是否是叶子节点
	IsLeaf() bool
}

func SummeryNodeInfo(root Node) (int, int) {
	testFileCnt := 0
	testCnt := 0
	for _, child := range root.GetChildren() {

		switch n := child.(type) {
		case *TestFileNode:
			testFileCnt += 1
			testCnt += n.TestCnt
		case *TestNode:
			if n.IsLeaf() {
				testCnt += 1
			} else {
				testCnt += n.TestCnt
			}
		case *SubTestNode:
			testCnt += 1
		case *DirNode:
			testFileCnt += n.TestFileCount
			testCnt += n.TestCount
		}
	}
	return testFileCnt, testCnt
}
