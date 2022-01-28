package cmd

import (
	"MaimuStreamer/global"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"text/tabwriter"
)

var tableWriter  *tabwriter.Writer
var version bool

var rootCmd = &cobra.Command{
	Use:   "MaimuStreamer",
	Short: "Maimu Streamer",
	Long:  `Multithreaded and multiplatform live streaming recorder.`,
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			color.Blue("Maimu Streamer\n\nVersion: %s\nRuntime: %s", global.Version, global.GetRuntime())
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&version, "version", "v", false, "get verison")
}

func Init() {
	cobra.CheckErr(rootCmd.Execute())
}
