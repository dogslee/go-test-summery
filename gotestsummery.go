package gotestsummery

import (
	"go-test-summery/node"
	"go-test-summery/output"
	"go-test-summery/parser"
)

// CreateMarkMapFromTestDir create mark map from test dir
func CreateMarkMapFromTestDir(testDir string, outPutDir string) {
	root := node.NewDIRNode("testmap")
	parser.WalfDirGroupFile(testDir, root)
	output.NodeToMarkMap(root, outPutDir)
}
