package cmd

import (
	"MaimuStreamer/database"
	"fmt"
	"github.com/spf13/cobra"
)

var listasksCmd = &cobra.Command{
	Use:   "listasks",
	Short: "List tasks in progress",
	Long:  `List tasks in progress`,
	Run: func(cmd *cobra.Command, args []string) {
		err := database.Init(dbFile)
		if err != nil {
			panic("Unable to open database file: " + err.Error())
		}
		tasks, err := database.GetAllTasks()
		if err != nil {
			panic("Unable to read all tasks: " + err.Error())
		}

		fmt.Fprintln(tableWriter, "[ID]\t[Platform]\t[Room ID]\t[Name]\t[Room Title]\t[Download Count]\t")
		for _, v := range *tasks {
			fmt.Fprintf(tableWriter, "%d\t%s\t%s\t%s\t%s\t%d\t\n", v.ID, v.Platform, v.RoomID, v.Name, v.RoomTitle, v.DLByteCount)
		}
		tableWriter.Flush()
		fmt.Println("\nDone.")
	},
}

func init() {
	rootCmd.AddCommand(listasksCmd)

	listasksCmd.PersistentFlags().StringVarP(&dbFile, "database", "d", "data.db", "database file path")

}
