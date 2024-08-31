package core

import (
	"os"
	"strings"
)

type OsArgs struct {
	Folder string
	Short  bool
	Delete string
	Help   bool
}

const (
	Folder   string = "-f"
	Short    string = "-s"
	Delete   string = "-d"
	Help     string = "-h"
	HelpFull string = "--help"
)

func (o *OsArgs) GetArgs() *OsArgs {
	args := os.Args[1:]
	if len(args) > 0 {
		for _, ele := range args {
			switch {
			case Includes(ele, Folder):
				o.Folder = strings.Split(ele, "=")[1]
			case Includes(ele, Short):
				o.Short = true
			case Includes(ele, Delete):
				o.Delete = strings.Split(ele, "=")[1]
			case Includes(ele, Help) || Includes(ele, HelpFull):
				o.Help = true
			default:
				return o
			}
		}
	}
	return o
}

func Includes(str, s string) bool {
	return strings.Contains(str, s)
}
