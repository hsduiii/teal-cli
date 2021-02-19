package cli

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/urfave/cli/v2"
)

//Run starts the CLI app
func Run() {
	app := cli.NewApp()
	app.Name = `
	
+++++++ ++++++      +     +             +++++  +       +
   +    +          + +    +             +      +       +
   +    ++++++    +   +   +             +      +       +
   +    +        +	   +  +             +      +       +
   +    ++++++  +   +++++ ++++++        +++++  ++++++  +

	`
	app.Usage = "@hsduiii"

	showFlags := []cli.Flag{
		&cli.StringFlag{
			Name:     "info",
			Value:    "cv",
			Required: true,
		},
	}

	openFlags := []cli.Flag{
		&cli.StringFlag{
			Name:     "app",
			Value:    "tealcloud",
			Required: true,
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:  "show",
			Usage: "--info cv",
			Flags: showFlags,
			Action: func(ctx *cli.Context) error {
				if ctx.NumFlags() == 0 {
					fmt.Println("A flag is required")
					return nil
				}
				if ctx.String("info") == "cv" {
					fmt.Println(`
-----------------------------------------------------------------------------------------------------------
                                            FULL STACK DEVELOPER

HECTOR ALONSO SALINAS GARCIA (hector_sag@outlook.com)

MONTE NUOVO #184 CUMBRES PROVENZA

CONTACT PHONE: 8120402661
-----------------------------------------------------------------------------------------------------------
PROFILE

	I'm a Full Stack Developer with knowledge in technical subjects such as Backend / Frontend Development, Deployments, Databases, Docker and Kubernetes.
	I like Linux environments, my favorite programming language is Golang and I'm aiming to become a DevOps Engineer and also getting the CKA / LFCS 
	certifications.

EDUCATION

	2014-2020
	SYSTEMS ADMINISTRATOR ENGINEER - FACULTAD DE INGENIERÍA MECÁNICA Y ELÉCTRICA

WORK EXPERIENCE

	-2020- FULL STACK DEVELOPER AT DALTUM SYSTEMS
	-2020-2020 – FULL STACK DEVELOPER AT INMAR WINSTON DATA (6 MONTHS)
	-2019-2020 – FULL STACK DEVELOPER AT OSIRIS HEALTHTECH SYSTEMS(1 year 2 months)
	-2018-2019 - BACKEND INTERN AT CONCÉNTRICO (1 year)
	-2016-2018 – INTERN AT NHP SOFT (2 years)

SKILLS

	-PROGRAMMING LANGUAGES: JAVASCRIPT/TYPESCRIPT, GOLANG, PHP, C#, PYTHON
	-BACKEND FRAMEWORKS: ECHO, KOA, FLASK, LARAVEL, SLIM
	-FRONTEND FRAMEWORKS: REACT, EMBERJS
	-PLATFORMS/ORCHESTRATION: DOCKER / KUBERNETES
	-DATABASES: POSTGRESQL, MYSQL, MSSQL
	-OS: LINUX, WINDOWS
	-INFRASTRUCTURES: HEROKU, GOOGLE CLOUD PLATFORM, IIS
	-VERSION CONTROL: GIT
	-CI/CD: BITBUCKET PIPELINES
	-CMS’S: WORDPRESS
	-LANGUAGES: SPANISH (NATIVE), ENGLISH(SEMI-ADVANCED).

PROJECTS
	
	TealCloud --> https://tealcloud.herokuapp.com/

WEBSITE
	
	https://hsduiii.github.io/portfolio/index.html
				`)
				}
				return nil
			},
		},
		{
			Name:  "open",
			Usage: "--app [tealcloud]",
			Flags: openFlags,
			Action: func(ctx *cli.Context) error {
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
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
