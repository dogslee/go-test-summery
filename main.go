package main

import (
	"fmt"
	"go-test-summery/node"
	"go-test-summery/output"
	"go-test-summery/parser"
	"os"
)

func main() {
	dir := os.Args[1]
	fmt.Print(dir)
	root := node.NewDIRNode("testmap")
	parser.WalfDirGroupFile(dir, root)
	// bytes, err := json.MarshalIndent(root, "", "  ")
	// if err != nil {
	// 	panic(err)
	// }
	// // fmt.Println(string(bytes))

	output.NodeToMarkMap(root, dir)
}
