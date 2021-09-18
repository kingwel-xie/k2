package cmd

import (
	"fmt"
	"github.com/kingwel-xie/k2/cmd/app"
	"github.com/kingwel-xie/k2/cmd/gen"
	"github.com/kingwel-xie/k2/cmd/version"
	"os"

	"github.com/kingwel-xie/k2/common/global"
	"github.com/kingwel-xie/k2/core/utils"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func tip() {
	usageStr := `欢迎使用 ` + `k2 `+global.Version + ` 可以使用 ` + `-h` + ` 查看命令`
	fmt.Printf("%s\n", usageStr)
}

func init() {
}

//Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func Init(name string, cmds... *cobra.Command) {
	rootCmd = &cobra.Command{
		Use:          name,
		Short:        name,
		SilenceUsage: true,
		Long:         name,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				tip()
				return fmt.Errorf(utils.Red("requires at least one arg"))
			}
			return nil
		},
		PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
		Run: func(cmd *cobra.Command, args []string) {
			tip()
		},
	}
	// TODO: DB migration
	//rootCmd.AddCommand(migrate.StartCmd)
	rootCmd.AddCommand(version.StartCmd)
	rootCmd.AddCommand(app.StartCmd)
	rootCmd.AddCommand(gen.StartCmd)
	//
	rootCmd.AddCommand(cmds...)
}
