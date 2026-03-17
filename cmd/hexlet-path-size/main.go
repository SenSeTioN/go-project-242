package main

import (
	"context"
	"fmt"
	"os"

	"code/internal/cli"
)

func main() {
	ctx := context.Background()
	cmd := cli.NewPathSizeCommand()

	if err := cmd.Run(ctx, os.Args); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
