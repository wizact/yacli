package yacli

import (
	"context"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"text/tabwriter"
)

// NewApplication returns a new Application instance with default values.
func NewApplication() *Application {
	return &Application{
		Name:        filepath.Base(os.Args[0]),
		Version:     "0.0.0",
		Description: "cli tool",
	}
}

// Application struct type.
// An application have multiple commands attached to it.
type Application struct {
	Name string

	Description string

	Version string

	Commands []Command
}


// Command interface with one or many flags.
type Command interface {
	Name() string

	HelpString() string

	Flags() *flag.FlagSet

	Run(ctx context.Context, args []string) error
}

// AddCommand adds a command to the application. It will do nothing if the command name already exists.
func (app *Application) AddCommand(c Command) {
	for _, command := range app.Commands {
		if command.Name() == c.Name() {
			return
		}
	}

	app.Commands = append(app.Commands, c)
}

// Run the command line integration pipeline.
func (app *Application) Run(ctx context.Context) {
	if len(os.Args) < 2 {
		logrus.Fatalf("Expected at least one sub command.")
	}

	commandName := os.Args[1]

	command, err := app.findCommand(commandName)
	if ae, ok := err.(*ArgError); ok {
		app.printCommandHelp(ae.Error())
		os.Exit(1)
	} else if err != nil {
		logrus.Fatalf(err.Error())
		os.Exit(1)
	}

	parseErr := command.Flags().Parse(os.Args[2:])
	if parseErr != nil {
		logrus.Fatalf(parseErr.Error())
		os.Exit(1)
	}

	runErr := command.Run(ctx, os.Args)

	if runErr != nil {
		logrus.Fatalf(runErr.Error())
	}
}

func (app *Application) printCommandHelp(errorMessage string) {

	if errorMessage != "" {
		fmt.Fprintln(os.Stderr, errorMessage)
		fmt.Fprintln(os.Stderr)
	}

	fmt.Fprintln(os.Stderr, "Supported commands:")
	fmt.Fprintln(os.Stderr)
	w := tabwriter.NewWriter(os.Stderr, 0, 4, 2, ' ', 0)
	for _, s := range app.Commands {
		fmt.Fprintf(w, "\t%s\t%s\n", s.Name(), s.HelpString())
	}
	w.Flush()
	fmt.Fprintln(os.Stderr)
}

func (app *Application) findCommand(name string) (Command, error) {
	for _, command := range app.Commands {
		if command.Name() == name {
			return command, nil
		}
	}

	return nil , &ArgError{argument: name}
}

// ArgError throws when the argument to cli does not satisfy the expectations
type ArgError struct {
	argument string
}

func (e *ArgError) Error() string {
	return fmt.Sprintf("command %s is not valid", e.argument)
}
