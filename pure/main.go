package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/prog-lang/pure/machine"
	"github.com/urfave/cli/v2"
)

var app = cli.App{
	Name:      "pure",
	Usage:     "Execute Pure bytecode",
	Flags:     []cli.Flag{debug},
	Args:      true,
	ArgsUsage: "<SOURCE>",
	Action:    run,
}

var debug = &cli.BoolFlag{
	Name:    "debug",
	Aliases: []string{"d"},
	Usage:   "run in debug mode",
	Action: func(ctx *cli.Context, b bool) error {
		log.SetLevel(log.DebugLevel)
		return nil
	},
}

func run(ctx *cli.Context) error {
	name := ctx.Args().First()
	if name == "" {
		log.Error("Did you forget to specify the SOURCE file?")
		return cli.ShowAppHelp(ctx)
	}

	src, err := machine.SourceFromFile(name)
	if err != nil {
		return err
	}

	cmd, err := src.Main()
	if err != nil {
		return err
	}

	cmd.Feed(machine.Unit{})
	return nil
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Error(err)
	}
}
