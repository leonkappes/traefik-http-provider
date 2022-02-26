package main

import (
	"log"
	"os"

	"github.com/leonkappes/go-traefik-daemon/internal/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "service",
				Aliases: []string{"s"},
				Usage:   "Perform action on a traefik service",
				Subcommands: []*cli.Command{
					{
						Name:   "info",
						Usage:  "Print json info of service",
						Action: commands.ServiceInfo,
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
