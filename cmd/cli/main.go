package main

import (
	"log/slog"
	"os"

	"github.com/happsie/trace2dia/internal"
	"github.com/happsie/trace2dia/internal/generator"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "t2d",
		Usage: "Transform traces to diagrams. Pass the path to the json trace file as argument",
		Action: func(c *cli.Context) error {
			tracePath := c.Args().First()
			trace, err := internal.LoadTraceFromPath(tracePath)
			if err != nil {
				return err
			}
			generationMethod := generator.NewPlantUML()
			return generationMethod.Generate(trace)
		},
	}
	if err := app.Run(os.Args); err != nil {
		slog.Error("could not run t2d", "error", err)
		os.Exit(1)
	}
}
