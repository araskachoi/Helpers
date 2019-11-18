package generate

import (
	"os"
	"fmt"
	"strconv"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
//	"path/filepath"
)

var (
	GenerateCommand = cli.Command{
		Name:        "generate",
		Usage:       "Command to generate ethereum accounts",
		Description: `Creation of public key, private key, and account address`,
		Action:      generateAccounts,
		Flags: []cli.Flag{
			ExportAsFileFlag,
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
	if ctx.Bool(ExportAsFileFlag.Name) == true {
                fmt.Println("exporting as file")
		err := ioutil.WriteFile("./accounts.txt", []byte(out), 0777)
                if err != nil {
                        log.Fatal("There was an error saving testnet id to file", err)
                }
        }
}
