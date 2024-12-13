package main

import (
	"size/core"
)

func main() {
	coreArgs := core.OsArgs{}
	cmdArgs := coreArgs.GetArgs()
	operation := core.NewOperation(cmdArgs)
	operation.Operate()
}
