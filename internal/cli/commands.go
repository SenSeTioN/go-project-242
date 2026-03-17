package cli

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func NewPathSizeCommand() *cli.Command {
	return &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "recursive size of directories",
			},
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit)",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			args := c.Args()
			if args.Len() < 1 {
				return fmt.Errorf("path argument is required")
			}

			path := args.First()
			human := c.Bool("human")
			all := c.Bool("all")
			recursive := c.Bool("recursive")

			size, err := GetSize(path, all, recursive)
			if err != nil {
				return err
			}

			formatted := FormatSize(size, human)
			fmt.Printf("%s\t%s\n", formatted, path)

			return nil
		},
	}
}
