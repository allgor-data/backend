package cmd

import (
	"github.com/urfave/cli/v2"
)

func Run(args []string) error {
	app := cli.NewApp()

	app.Name = "allgor"
	app.Version = "0.0.1"
	app.Copyright = "(c) 2023 Allgor"

	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:    "playground",
			Aliases: []string{"t"},
			EnvVars: []string{"ALLGOR_PLAYGROUND"},
			Usage:   "Enable GraphQL playground",
		},
		&cli.StringFlag{
			Name:    "port",
			Aliases: []string{"p"},
			EnvVars: []string{"ALLGOR_PORT"},
			Usage:   "GraphQL server port",
			Value:   "8080",
		},
	}

	app.Action = func(c *cli.Context) error {
		return startServer(c.String("port"), c.Bool("playground"))
	}

	return app.Run(args)
}
