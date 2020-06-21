package main

import (
	"context"
	"github.com/wizact/yacli"
)

func main() {
	app := yacli.NewApplication()

	app.Name = "Test cli"
	app.Description = "Test cli app"
	app.Version = "0.0.1"

	app.AddCommand(&FooCommand{})


	type memberIDCtxKey string
	ctx := context.WithValue(context.Background(), memberIDCtxKey("memberId"), "123")

	app.Run(ctx)
}
