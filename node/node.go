package node

type Node interface {
	// 获取子节点list
	GetChildren() []Node

	// 获取节点名称
	GetName() string

	// 添加子节点
	AddChild(node Node)

	// 是否是叶子节点
	IsLeaf() bool
}
