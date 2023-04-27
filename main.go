package main

import (
	"os"

	"github.com/dogslee/go-test-summery/node"
	"github.com/dogslee/go-test-summery/output"
	"github.com/dogslee/go-test-summery/parser"
)

var (
	DefaultTestMapName = "testmap"
)

func main() {
	testDir := "./"
	outPutDir := "./"
	if len(os.Args) > 1 {
		testDir = os.Args[1]
	}

	if len(os.Args) > 2 {
		outPutDir = os.Args[2]
	}

	if len(os.Args) > 3 {
		DefaultTestMapName = os.Args[3]
	}

	root := node.NewDIRNode(nil, DefaultTestMapName)
	parser.WalfDirGroupFile(testDir, root)
	root.TestFailCount, root.TestCount = node.SummeryNodeInfo(root)
	output.NodeToMarkMap(root, outPutDir)
}
