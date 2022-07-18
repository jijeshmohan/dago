package commands

import (
	"context"
	"fmt"

	"github.com/jijeshmohan/dago/pkg/config"
	"github.com/jijeshmohan/dago/pkg/generator"
	"github.com/jijeshmohan/dago/pkg/xlogger"
	"github.com/urfave/cli/v2"
)

func Generate() *cli.Command {
	return &cli.Command{
		Name:      "generate",
		Aliases:   []string{"g", "gen"},
		Usage:     "generate from a template",
		UsageText: "generate [template-name] [path]",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "~/.config/dago",
				Usage:   "Configuration folder path",
			},
		},
		Action: func(c *cli.Context) error {
			configPath := c.String("config")
			conf, err := config.Load(configPath)
			if err != nil {
				return err
			}

			if c.NArg() < 2 {
				fmt.Println(c.Command.UsageText)
				return cli.Exit("missing arguments, please check usage", 1)
			}

			ctx := context.Background()
			logger := xlogger.NewColorLogger(xlogger.INFO, true)

			g, err := generator.NewGenerator(conf, logger)
			if err != nil {
				return err
			}

			return g.Generate(ctx, c.Args().First(), c.Args().Get(1))
		},
	}
}
