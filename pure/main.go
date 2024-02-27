package main

import (
	"os"
	"os/exec"

	"github.com/charmbracelet/log"
	"github.com/prog-lang/pure/machine"
	"github.com/urfave/cli/v2"
)

var app = cli.App{
	Name:     "pure",
	Usage:    "Language for purely functional microservices",
	Flags:    []cli.Flag{debug},
	Action:   cli.ShowAppHelp,
	Commands: []*cli.Command{com, exe, run},
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

var (
	com = &cli.Command{
		Name:    "com",
		Aliases: []string{"c"},
		Usage:   "Compile Pure to bytecode",
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "Path to output file",
				Value:   "main.pure.exe",
			},
		},
		Args:      true,
		ArgsUsage: "<SOURCE>",
		Action: func(ctx *cli.Context) error {
			name := ctx.Args().First()
			if name == "" {
				log.Error("Did you forget to specify the SOURCE file?")
				return cli.ShowAppHelp(ctx)
			}
			return purec(ctx, name, ctx.Path("output"))
		},
	}

	exe = &cli.Command{
		Name:      "exe",
		Aliases:   []string{"e"},
		Usage:     "Execute Pure from bytecode",
		Args:      true,
		ArgsUsage: "<SOURCE>",
		Action: func(ctx *cli.Context) error {
			name := ctx.Args().First()
			if name == "" {
				log.Error("Did you forget to specify the SOURCE file?")
				return cli.ShowAppHelp(ctx)
			}
			return execute(name)
		},
	}

	run = &cli.Command{
		Name:      "run",
		Aliases:   []string{"r"},
		Usage:     "Instantly run Pure from source",
		Args:      true,
		ArgsUsage: "<SOURCE>",
		Action: func(ctx *cli.Context) error {
			const output = "tmp.pure.exe"
			name := ctx.Args().First()
			if name == "" {
				log.Error("Did you forget to specify the SOURCE file?")
				return cli.ShowAppHelp(ctx)
			}
			if err := purec(ctx, name, output); err != nil {
				return err
			}
			defer os.Remove(output)
			return execute(output)
		},
	}
)

func purec(ctx *cli.Context, name, output string) error {
	cmd := exec.Command("purec", name, "--output", output)
	cmd.Stdin = ctx.App.Reader
	cmd.Stdout = ctx.App.Writer
	cmd.Stderr = ctx.App.ErrWriter
	return cmd.Run()
}

func execute(name string) error {
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
