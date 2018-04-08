package cmd

import (
	"github.com/aestek/dotfiler/sync"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(remoteCmd)
}

var remoteCmd = &cobra.Command{
	Use:   "remote <git url>",
	Short: "Init dotfiler",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := getConfig()
		must(sync.SetRemote(cfg.Workdir, args[0]))
	},
}
