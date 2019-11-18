package generate

import (
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"path/filepath"
)

var (
	GenerateCommand = cli.Command{
		Name:        "generate <number_of_nodes>",
		Usage:       "Command to generate ethereum accounts",
		Description: `Creation of public key, private key, and account address`,
		Action:      generateAccounts,
		Flags: []cli.Flag{
		},
	}
)

func generateAccounts(nodes int) {
	println(nodes)
}
