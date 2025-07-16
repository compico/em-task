package main

import (
	"context"
	"github.com/compico/em-task/cmd/commands"
	_ "github.com/compico/em-task/cmd/commands/http"
	"github.com/urfave/cli/v3"
	"log/slog"
	"os"
	"strings"
)

func main() {
	app := &cli.Command{
		Commands: commands.Commands,
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		slog.Error(
			"error on start app",
			"command", strings.Join(os.Args[1:], " "),
			"error", err,
		)
	}
}
