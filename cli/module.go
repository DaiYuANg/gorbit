package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// NewCLIModule 返回 fx.Module
func NewCLIModule(opts ...CLIOption) fx.Option {
	options := defaultCLIOptions()
	for _, o := range opts {
		o(options)
	}

	return fx.Module("cli_module",
		fx.Provide(func() *cobra.Command {
			root := &cobra.Command{
				Use:     options.RootName,
				Version: options.Version,
				PersistentPreRun: func(cmd *cobra.Command, args []string) {
					if options.PreRunHook != nil {
						options.PreRunHook(cmd, args)
					}
				},
			}

			for _, cmd := range options.Commands {
				root.AddCommand(cmd)
			}

			return root
		}),
		fx.Invoke(func(cmd *cobra.Command) {
			// 注意这里不直接调用 Execute()，可以在 fx.App 完成生命周期后再调用
			fmt.Println("CLI module ready, use cmd.Execute() to run CLI")
		}),
	)
}
