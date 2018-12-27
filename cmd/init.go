package cmd

import (
	"fmt"

	"github.com/miseyu/go-blockchain-bootstrap/lib"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize go blockchain bootstrap settings",
	Run: func(cmd *cobra.Command, args []string) {
		if err := lib.Init(); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
