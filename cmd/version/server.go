package version

import (
	"fmt"
	"github.com/kingwel-xie/k2/common/global"
	"github.com/spf13/cobra"
)

var (
	StartCmd = &cobra.Command{
		Use:     "version",
		Short:   "Show version info",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func run() error {
	fmt.Println(global.Version)
	return nil
}
