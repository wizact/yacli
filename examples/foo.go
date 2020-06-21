package main

import (
	"context"
	"flag"
	"fmt"
)

// FooCommand is foo command
type FooCommand struct {
	bar string
	baz bool
}

// Flags returns the flag sets
func (fc *FooCommand) Flags() *flag.FlagSet {
	f := &flag.FlagSet{}

	f.StringVar(&fc.bar, "bar", "barDefault", "Bar name")
	f.BoolVar(&fc.baz, "baz", false, "Baz value")

	return f
}

func (fc *FooCommand) getBar() string {
	return fc.bar
}

func (fc *FooCommand) getBaz() bool {
	return fc.baz
}

// Name of the command
func (fc *FooCommand) Name() string {
	return "foo"
}

// HelpString is the string shown as usage
func (fc *FooCommand) HelpString() string {
	return "foo"
}

// Run a command
func (fc *FooCommand) Run(ctx context.Context, args []string) error {
	fmt.Println(fc.getBar())
	fmt.Println(fc.getBaz())
	return nil
}
