package cmd

import (
	"MaimuStreamer/config"
	"MaimuStreamer/database"
	"MaimuStreamer/task"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

var (
	configFile string
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start the task pool",
	Long:  `Start the task pool`,
	Run: func(cmd *cobra.Command, args []string) {
		err := config.Init(configFile)
		if err != nil {
			panic("Unable to read config file: " + err.Error())
		}
		err = database.Init(config.Cfg.SqliteDb)
		if err != nil {
			panic("Unable to open database file: " + err.Error())
		}

		task.Init()
	},
}

func init() {
	tableWriter = tabwriter.NewWriter(os.Stdout, 1, 1, 2, ' ', 0)
	rootCmd.AddCommand(runCmd)

	runCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config.json", "config file")
}
