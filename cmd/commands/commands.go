package commands

import "github.com/urfave/cli/v3"

var Commands []*cli.Command

func RegisterCommand(c *cli.Command) {
	Commands = append(Commands, c)
}
