package main

import (
	"fmt"
	"os"

	"github.com/Beyond-simplechain/erc20msig/utils"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

func init() {
	app.Commands = []cli.Command{
		utils.SignCommand,
		//utils.SendCommand,
	}

}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
