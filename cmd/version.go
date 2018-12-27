package cmd

import (
	"fmt"
	"runtime"

	"github.com/miseyu/go-blockchain-bootstrap/lib"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s (built with %s)\n", lib.Version, runtime.Version())
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
