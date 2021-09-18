// +build admin

package cmd

import "github.com/kingwel-xie/k2/cmd/admin"

func init() {
	rootCmd.AddCommand(admin.StartCmd)
}
