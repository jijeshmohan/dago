package commands

import (
	"fmt"
	"os"

	"github.com/jijeshmohan/dago/pkg/config"
	"github.com/jijeshmohan/dago/pkg/templates"
	"github.com/jijeshmohan/dago/pkg/xfilesystem"
	"github.com/urfave/cli/v2"
)

func Template() *cli.Command {
	return &cli.Command{
		Name:    "template",
		Aliases: []string{"t"},
		Usage:   "dago template management",
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
				Usage: "create template",
				Action: func(c *cli.Context) error {
					configPath := c.String("config")
					conf, err := config.Load(configPath)
					if err != nil {
						return err
					}

					if _, err := templates.CreateTemplate(c.Args().First(), conf.TemplatesPath); err != nil {
						return err
					}

					fmt.Println("Template created successfully")
					return nil
				},
			},
			{
				Name:  "validate",
				Usage: "validate a given template",
				Action: func(c *cli.Context) error {
					panic("TODO")
				},
			},
			{
				Name:  "list",
				Usage: "list all available templates",
				Action: func(c *cli.Context) error {
					configPath := c.String("config")
					conf, err := config.Load(configPath)
					if err != nil {
						return err
					}

					repo, err := templates.NewFSRepository(xfilesystem.NewFileSystem(conf.TemplatesPath, os.DirFS(conf.TemplatesPath)))
					if err != nil {
						return err
					}

					for _, templateName := range repo.GetAllTemplateNames() {
						fmt.Println(templateName)
					}

					return nil
				},
			},
		},
	}
}
