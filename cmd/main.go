package main

import (
	"context"
	"github.com/compico/em-task/cmd/commands"
	_ "github.com/compico/em-task/cmd/commands/http"
	_ "github.com/compico/em-task/cmd/commands/migrate"
	"github.com/urfave/cli/v3"
	"log/slog"
	"os"
	"strings"
)

func main() {
	app := &cli.Command{
		Commands: commands.Commands,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE`",
				Value:   "configs/config.yaml",
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		slog.Error(
			"error on start app",
			"command", strings.Join(os.Args[1:], " "),
			"error", err,
		)
	}
}
