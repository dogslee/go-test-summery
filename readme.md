# go-test-summery

将所有单元测试内容汇总到一个思维导图中

## 用法

```go
package main

import (
	testSummery "github.com/dogslee/go-test-summery"
)

func main() {
	// testdir is the directory of test files
	// output is the directory of output files
	testSummery.CreateMarkMapFromTestDir("./testdir", "./output")
}

```
