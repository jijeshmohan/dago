package commands

import (
	"fmt"

	"github.com/jijeshmohan/dago/pkg/config"
	"github.com/urfave/cli/v2"
)

func Config() *cli.Command {
	return &cli.Command{
		Name:  "config",
		Usage: "dago configuration",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "~/.config/dago",
				Usage:   "Configuration folder path",
			},
		},
		Subcommands: []*cli.Command{
			{
				Name:  "new",
				Usage: "create new configuration",
				Action: func(c *cli.Context) error {
					configPath := c.String("config")
					err := config.Create(configPath)
					if err != nil {
						return err
					}

					fmt.Println("Configuration created successfully")
					return nil
				},
			},
			{
				Name:  "show",
				Usage: "show current configuration",
				Action: func(c *cli.Context) error {
					configPath := c.String("config")

					conf, err := config.Load(configPath)
					if err != nil {
						return err
					}

					fmt.Printf("\nUsing configuration from path %s\n", configPath)
					fmt.Println("\n---")
					fmt.Printf("%s", conf.String())
					return nil
				},
			},
		},
	}
}
