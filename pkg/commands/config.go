package commands

import (
	"github.com/urfave/cli/v2"
)

func Config() *cli.Command {
	return &cli.Command{
		Name:  "config",
		Usage: "dago configuration",
		Subcommands: []*cli.Command{
			{
				Name:  "new",
				Usage: "create new configuration",
			},
		},
	}
}
