package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jijeshmohan/dago/pkg/commands"
	"github.com/urfave/cli/v2"
)

var (
	version = "development"
)

func main() {
	app := newCliApp()

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

func newCliApp() *cli.App {
	app := cli.NewApp()
	app.Name = "Dago"
	app.HelpName = filepath.Base(os.Args[0])
	app.Usage = "Code generation using template"
	app.Version = version
	app.Copyright = "(c) Jijesh Mohan"
	app.EnableBashCompletion = true

	registerCommands(app)

	return app
}

func registerCommands(app *cli.App) {
	app.Commands = []*cli.Command{
		commands.Config(),
	}
}
