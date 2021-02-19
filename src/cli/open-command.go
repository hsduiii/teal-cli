package cli

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/urfave/cli/v2"
)

//OpenCommand variable
var OpenCommand *cli.Command = &cli.Command{
	Name:    "open",
	Usage:   "Open one of my personal projects",
	Aliases: []string{"o"},
	Flags:   GetOpenFlags(),
	Action: func(ctx *cli.Context) error {
		ctx.NArg()
		if ctx.NumFlags() == 0 {

			fmt.Println("A flag is required")
			return nil

		}

		if ctx.String("app") == "tealcloud" {

			url := "https://tealcloud.herokuapp.com/"

			fmt.Println(`Opening TealCloud...`)

			var err error

			switch runtime.GOOS {
			case "linux":
				err = exec.Command("xdg-open", url).Start()
			case "windows":
				err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
			case "darwin":
				err = exec.Command("open", url).Start()
			default:
				err = fmt.Errorf("unsupported platform")
			}

			if err != nil {
				log.Fatal("Could not open TealCloud")
			}
		}

		return nil
	},
}
