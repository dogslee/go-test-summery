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

	depthFirstSearch(file, root, 1)

}

// 深度遍历node.Node 输出深度
func depthFirstSearch(file *os.File, root node.Node, depth int) {
	switch n := root.(type) {
	case *node.DirNode:
		file.WriteString("\n")
		file.WriteString(fmt.Sprintf("%s %s\n", depthLevel(depth), n.GetName()))
		file.WriteString("\n")
	case *node.TestFileNode:
		file.WriteString("\n")
		file.WriteString(fmt.Sprintf("%s %s\n", depthLevel(depth), n.GetName()))
		file.WriteString("\n")
	case *node.TestNode:
		if n.IsLeaf() {
			file.WriteString(fmt.Sprintf("- %s\n", n.GetName()))
		} else {
			file.WriteString("\n")
			file.WriteString(fmt.Sprintf("%s %s\n", depthLevel(depth), n.GetName()))
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
