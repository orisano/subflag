package subflag

import (
	"errors"
	"flag"
	"fmt"
)

type Command interface {
	FlagSet() *flag.FlagSet
	Run(args []string) error
}

var ErrInvalidArguments = errors.New("invalid arguments")

func SubCommand(args []string, commands []Command) error {
	if len(args) == 0 {
		return fmt.Errorf("subcommand is required")
	}
	subCommand := args[0]

	for _, command := range commands {
		flagSet := command.FlagSet()
		if flagSet.Name() != subCommand {
			continue
		}
		if err := flagSet.Parse(args[1:]); err != nil {
			return err
		}
		err := command.Run(flagSet.Args())
		if err == ErrInvalidArguments && flagSet.Usage != nil {
			flagSet.Usage()
		}
		return err
	}
	return fmt.Errorf("subcommand was not found: %q", subCommand)
}
