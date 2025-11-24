package cli

import "github.com/spf13/cobra"

// 功能型 Option
type CLIOption func(*cliOptions)

type cliOptions struct {
	RootName   string
	Version    string
	Commands   []*cobra.Command
	PreRunHook func(cmd *cobra.Command, args []string)
}

// 默认值
func defaultCLIOptions() *cliOptions {
	return &cliOptions{
		RootName:   "app",
		Version:    "v0.1.0",
		Commands:   []*cobra.Command{},
		PreRunHook: nil,
	}
}

// Option helper
func WithRootName(name string) CLIOption {
	return func(o *cliOptions) { o.RootName = name }
}

func WithVersion(version string) CLIOption {
	return func(o *cliOptions) { o.Version = version }
}

func WithCommand(cmd *cobra.Command) CLIOption {
	return func(o *cliOptions) { o.Commands = append(o.Commands, cmd) }
}

func WithPreRunHook(hook func(cmd *cobra.Command, args []string)) CLIOption {
	return func(o *cliOptions) { o.PreRunHook = hook }
}
