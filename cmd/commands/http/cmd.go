package http

import (
	"context"
	"github.com/compico/em-task/cmd/commands"
	"github.com/urfave/cli/v3"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	commands.RegisterCommand(
		&cli.Command{
			Name: "http",
			Action: func(ctx context.Context, cmd *cli.Command) error {
				app, dbCloseFunc, err := InitializeApp(
					ctx,
					cmd.String("config"),
				)
				if err != nil {
					return err
				}
				app.logger.Info("Application initialized")

				go func() {
					app.logger.Info("Starting web server")
					err := app.server.Start()
					app.logger.ErrorContext(ctx, "error on start server", "error", err)
				}()

				sigCh := make(chan os.Signal, 1)
				signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
				<-sigCh

				app.logger.Info("Closing db connection")
				dbCloseFunc()

				app.logger.Info("Shutting down web server")
				return app.server.Stop(ctx)
			},
		},
	)
}
