package http

import (
	"context"
	"github.com/compico/em-task/cmd/commands"
	"os"
	"os/signal"
	"syscall"

	"github.com/urfave/cli/v3"
)

func init() {
	commands.RegisterCommand(
		&cli.Command{
			Name: "http",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "config",
					Aliases: []string{"c"},
					Usage:   "Load configuration from `FILE`",
					Value:   "configs/config.yaml",
				},
			},
			Action: func(ctx context.Context, cmd *cli.Command) error {
				app, err := InitializeApp(
					ctx,
					cmd.String("config"),
				)
				if err != nil {
					return err
				}

				go func() {
					err := app.server.Start()
					app.logger.ErrorContext(ctx, "error on start server", "error", err)
				}()

				sigCh := make(chan os.Signal, 1)
				signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
				<-sigCh

				//err = app.db.Close(ctx)
				//if err != nil {
				//	return err
				//}
				return app.server.Stop(ctx)
			},
		},
	)
}
