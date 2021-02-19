package cli

import "github.com/urfave/cli/v2"

//GetShowFlags for the show command
func GetShowFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "info",
			Value:    "cv",
			Required: true,
		},
	}
}
