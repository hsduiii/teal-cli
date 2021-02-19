package cli

import "github.com/urfave/cli/v2"

//GetOpenFlags for the open command
func GetOpenFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "app",
			Value:    "tealcloud",
			Required: true,
		},
	}
}
