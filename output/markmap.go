package output

import (
	"fmt"
	"os"

	"github.com/dogslee/go-test-summery/node"
)

// 将node.Node 转化为markdown文本格式
func NodeToMarkMap(root node.Node, dir string) {
	// 打开一个文件写入数据
	file, err := os.OpenFile(dir+"/testmap.md", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	depthFirstSearchCount(root)

	depthFirstSearch(file, root, 1)

}

// 深度遍历统计测试case信息
func depthFirstSearchCount(root node.Node) {
	if root == nil {
		return
	}
	for _, child := range root.GetChildren() {
		depthFirstSearchCount(child)
	}
	parent := root.GetParent()
	if parent == nil {
		return
	}
	fn, tn := node.SummeryNodeInfo(parent)
	switch p := parent.(type) {
	case *node.SubTestNode:
	case *node.TestNode:
		if p.IsLeaf() {
			p.TestCnt = 1
		} else {
			p.TestCnt = tn
		}
	case *node.TestFileNode:
		p.TestCnt = tn
	case *node.DirNode:
		p.TestCount = tn
		p.TestFileCount = fn
	}
}

// 深度遍历node.Node 输出深度
func depthFirstSearch(file *os.File, root node.Node, depth int) {
	switch n := root.(type) {
	case *node.DirNode:
		file.WriteString("\n")
		file.WriteString(fmt.Sprintf("%s %s **文件数:%d, 用例数:%d**\n", depthLevel(depth), n.GetName(), n.TestFileCount, n.TestCount))
		file.WriteString("\n")
	case *node.TestFileNode:
		file.WriteString("\n")
		file.WriteString(fmt.Sprintf("%s %s **用例数:%d**\n", depthLevel(depth), n.GetName(), n.TestCnt))
		file.WriteString("\n")
	case *node.TestNode:
		if n.IsLeaf() {
			file.WriteString(fmt.Sprintf("%s %s\n", depthLevel(depth), n.GetName()))
		} else {
			file.WriteString("\n")
			file.WriteString(fmt.Sprintf("%s %s**用例数:%d**\n", depthLevel(depth), n.GetName(), n.TestCnt))
			file.WriteString("\n")
		}

	case *node.SubTestNode:
		file.WriteString(fmt.Sprintf("- %s\n", n.GetName()))
	}
	if root.IsLeaf() {
		return
	}
	for _, child := range root.GetChildren() {
		depthFirstSearch(file, child, depth+1)
	}
}

func depthLevel(depth int) string {
	ret := ""
	for i := 0; i < depth; i++ {
		ret += "#"
	}
	return ret
}
