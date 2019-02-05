package subflag

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

type Command interface {
	FlagSet() *flag.FlagSet
	Run(args []string) error
}

// ErrInvalidArguments is ...
// Deprecated: instead of use flag.ErrHelp
var ErrInvalidArguments = errors.New("invalid arguments (deprecated)")

func showSubCommands(commands []Command) string {
	var names []string
	for _, command := range commands {
		names = append(names, " - "+command.FlagSet().Name())
	}
	return "Available SubCommands:\n" + strings.Join(names, "\n")
}

func SubCommand(args []string, commands []Command) error {
	if len(args) == 0 {
		return fmt.Errorf("subcommand is required: \n%s", showSubCommands(commands))
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
		showUsage := err == flag.ErrHelp || err == ErrInvalidArguments // for backward compatibility
		if showUsage && flagSet.Usage != nil {
			flagSet.Usage()
		}
		return err
	}
	return fmt.Errorf("subcommand is not found: %q\n%s", subCommand, showSubCommands(commands))
}
