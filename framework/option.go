package framework

import (
	"context"

	"github.com/samber/do/v2"
	"github.com/spf13/cobra"
)

type Options struct {
	Name         string
	Description  string
	ConfigLoader func() (Config, error)
	RootRun      func(ctx context.Context, i *do.RootScope) error
	SubCommands  []*cobra.Command
}

type Option func(*Options)

func defaultOptions() Options {
	return Options{
		Name:        "app",
		Description: "A modular Go application framework",
	}
}

func WithName(name string) Option {
	return func(o *Options) { o.Name = name }
}

func WithDescription(desc string) Option {
	return func(o *Options) { o.Description = desc }
}

func WithConfigLoader(loader func() (Config, error)) Option {
	return func(o *Options) { o.ConfigLoader = loader }
}

func WithRootCommand(fn func(ctx context.Context, i *do.RootScope) error) Option {
	return func(o *Options) { o.RootRun = fn }
}

func WithSubCommand(cmd *cobra.Command) Option {
	return func(o *Options) { o.SubCommands = append(o.SubCommands, cmd) }
}
