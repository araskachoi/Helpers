package main

import (
	"os"
	"log"
	"github.com/urfave/cli"
	"github.com/araskachoi/Helpers/ethAccGen/generate"
)

func main() {
	app := cli.NewApp()
	app.Version = "1.0.0"
	app.Name = "ethAccGen"
	app.Usage = "Generate Ethereum Accounts"
	app.Commands = []cli.Command{
		generate.GenerateCommand,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
