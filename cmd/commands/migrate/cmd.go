package migrate

import (
	"context"
	"github.com/compico/em-task/cmd/commands"
	"github.com/urfave/cli/v3"
)

func init() {
	commands.RegisterCommand(
		&cli.Command{
			Name:     "migrate",
			Usage:    "Up migrations",
			Commands: []*cli.Command{},
			Action: func(ctx context.Context, cmd *cli.Command) error {
				m, err := InitializeMigrator(
					ctx,
					cmd.String("config"),
				)
				if err != nil {
					return err
				}

				return m.migrate.Up()
			},
		},
	)
}
