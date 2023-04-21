package gotestsummery

import (
	"github.com/dogslee/go-test-summery/node"
	"github.com/dogslee/go-test-summery/output"
	"github.com/dogslee/go-test-summery/parser"
)

// CreateMarkMapFromTestDir create mark map from test dir
func CreateMarkMapFromTestDir(testDir string, outPutDir string) {
	root := node.NewDIRNode("testmap")
	parser.WalfDirGroupFile(testDir, root)
	output.NodeToMarkMap(root, outPutDir)
}
