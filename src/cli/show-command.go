package cli

import (
	"fmt"

	"bitbucket.com/hsduiii/teal-cli/src/utils"
	"github.com/urfave/cli/v2"
)

//ShowCommand variable
var ShowCommand *cli.Command = &cli.Command{

	Name:    "show",
	Usage:   "Show information about me",
	Aliases: []string{"s"},
	Flags:   GetShowFlags(),
	Action: func(ctx *cli.Context) error {

		if ctx.NumFlags() == 0 {
			fmt.Println("A flag is required")
			return nil
		}

		if ctx.String("info") == "cv" {
			fmt.Println(utils.GetCV())
		}

		return nil

	},
}
