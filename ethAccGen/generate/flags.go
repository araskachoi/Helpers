package generate

import (
	"github.com/urfave/cli"
)

var (
	ExportAsFileFlag = cli.BoolFlag{
		Name:  "file",
		Usage: "Export output as files",
//		Value: false,
	}
)
