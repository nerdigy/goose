package cli

import (
	"context"
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/peterbourgon/ff/v3/ffcli"
)

func newUpToCmd(root *rootConfig) *ffcli.Command {
	fs := flag.NewFlagSet("goose up-to", flag.ExitOnError)
	root.registerFlags(fs)

	return &ffcli.Command{
		Name:    "up-to",
		FlagSet: fs,
		Exec:    execUpToCmd(root),
	}
}

func execUpToCmd(root *rootConfig) func(ctx context.Context, args []string) error {
	return func(ctx context.Context, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("missing required argument: version")
		}
		version, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid version: %s, must be an integer", args[0])
		}
		provider, err := newGooseProvider(root)
		if err != nil {
			return err
		}
		now := time.Now()
		results, err := provider.UpTo(ctx, version)
		if err != nil {
			return err
		}
		return printMigrationResult(
			results,
			time.Since(now),
			root.useJSON,
		)
	}
}
