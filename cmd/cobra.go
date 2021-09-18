package cmd

import (
	"errors"
	"fmt"
	"github.com/kingwel-xie/k2/cmd/app"
	"github.com/kingwel-xie/k2/cmd/gen"
	"github.com/kingwel-xie/k2/common/global"
	"github.com/kingwel-xie/k2/core/utils"
	"os"

	"github.com/spf13/cobra"

	"github.com/kingwel-xie/k2/cmd/migrate"
	"github.com/kingwel-xie/k2/cmd/version"
)

var rootCmd = &cobra.Command{
	Use:          "github.com/kingwel-xie/k2",
	Short:        "github.com/kingwel-xie/k2",
	SilenceUsage: true,
	Long:         `k2`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New(utils.Red("requires at least one arg"))
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `欢迎使用 ` + utils.Green(`k2 `+global.Version) + ` 可以使用 ` + utils.Red(`-h`) + ` 查看命令`
	fmt.Printf("%s\n", usageStr)
}

func init() {
	rootCmd.AddCommand(migrate.StartCmd)
	rootCmd.AddCommand(version.StartCmd)
	rootCmd.AddCommand(app.StartCmd)
	rootCmd.AddCommand(gen.StartCmd)
}

//Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
