package generate

import (
	"os"
	"fmt"
	"strconv"
	"github.com/urfave/cli"
)

var (
	GenerateCommand = cli.Command{
		Name:        "generate",
		Usage:       "Command to generate ethereum accounts",
		Description: `Creation of public key, private key, and account address`,
		Action:      generateAccounts,
		Flags: []cli.Flag{
		},
	}
)

func generateAccounts(ctx *cli.Context) {
	nodes, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
	}
	acc, err := GenerateAccounts(nodes)
	if err != nil {
		fmt.Println(err)
	}
	out, err := Export(acc)
	if err != nil {
                fmt.Println(err)
        }
	fmt.Println(out)
}
