package cli

import (
	"log"
	"os"

	"bitbucket.com/hsduiii/teal-cli/src/utils"
	"github.com/urfave/cli/v2"
)

//Run starts the CLI app
func Run() {
	app := cli.NewApp()

	app.Name = utils.GetHeader()

	app.Version = "1.0.0"

	app.Description = "Personal CLI made in Golang"

	app.Usage = "@hsduiii"

	app.Commands = []*cli.Command{
		ShowCommand,
		OpenCommand,
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
