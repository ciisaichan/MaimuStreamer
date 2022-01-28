package cmd

import (
	"MaimuStreamer/database"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	dbFile       string
	dbId         uint64
	roomPlatform string
	roomId       string
	roomName     string
)

var listroomsCmd = &cobra.Command{
	Use:   "listrooms",
	Short: "Show all live rooms",
	Long:  `Show all live rooms`,
	Run: func(cmd *cobra.Command, args []string) {
		err := database.Init(dbFile)
		if err != nil {
			panic("Unable to open database file: " + err.Error())
		}
		rooms, err := database.GetAllRooms()
		if err != nil {
			panic("Unable to read all rooms: " + err.Error())
		}

		fmt.Fprintln(tableWriter, "[ID]\t[Platform]\t[Room ID]\t[Name]\t")
		for _, v := range *rooms {
			fmt.Fprintf(tableWriter, "%d\t%s\t%s\t%s\t\n", v.ID, v.Platform, v.RoomID, v.Name)
		}
		tableWriter.Flush()
		fmt.Println("\nDone.")
	},
}

var addroomCmd = &cobra.Command{
	Use:   "addroom",
	Short: "Add live room",
	Long:  `Add live room`,
	Run: func(cmd *cobra.Command, args []string) {
		err := database.Init(dbFile)
		if err != nil {
			panic("Unable to open database file: " + err.Error())
		}
		room, err := database.AddLiveRoom(roomPlatform, roomId, roomName)
		if err != nil {
			panic("Unable to add room: " + err.Error())
		}

		fmt.Fprintln(tableWriter, "[ID]\t[Platform]\t[Room ID]\t[Name]\t")
		fmt.Fprintf(tableWriter, "%d\t%s\t%s\t%s\t\n", room.ID, room.Platform, room.RoomID, room.Name)
		tableWriter.Flush()
		fmt.Println("\nDone.")
	},
}

var delroomCmd = &cobra.Command{
	Use:   "delroom",
	Short: "Delete live room",
	Long:  `Delete live room`,
	Run: func(cmd *cobra.Command, args []string) {
		err := database.Init(dbFile)
		if err != nil {
			panic("Unable to open database file: " + err.Error())
		}
		err = database.DelLiveRoom(dbId)
		if err != nil {
			panic("Unable to delete room: " + err.Error())
		}
		fmt.Println("\nDone.")
	},
}

var editroomCmd = &cobra.Command{
	Use:   "editroom",
	Short: "Edit live room",
	Long:  `Edit live room`,
	Run: func(cmd *cobra.Command, args []string) {
		err := database.Init(dbFile)
		if err != nil {
			panic("Unable to open database file: " + err.Error())
		}
		room, err := database.UpdateLiveRoom(dbId, roomPlatform, roomId, roomName)
		if err != nil {
			panic("Unable to delete room: " + err.Error())
		}

		fmt.Fprintln(tableWriter, "[ID]\t[Platform]\t[Room ID]\t[Name]\t")
		fmt.Fprintf(tableWriter, "%d\t%s\t%s\t%s\t\n", room.ID, room.Platform, room.RoomID, room.Name)
		tableWriter.Flush()
		fmt.Println("\nDone.")
	},
}

func init() {
	rootCmd.AddCommand(listroomsCmd)
	rootCmd.AddCommand(addroomCmd)
	rootCmd.AddCommand(delroomCmd)
	rootCmd.AddCommand(editroomCmd)

	listroomsCmd.PersistentFlags().StringVarP(&dbFile, "database", "d", "data.db", "database file path")

	addroomCmd.PersistentFlags().StringVarP(&dbFile, "database", "d", "data.db", "database file path")
	addroomCmd.PersistentFlags().StringVarP(&roomPlatform, "platform", "p", "", "Room platform")
	addroomCmd.PersistentFlags().StringVarP(&roomId, "roomid", "r", "", "Room id")
	addroomCmd.PersistentFlags().StringVarP(&roomName, "name", "n", "", "Room name")

	delroomCmd.PersistentFlags().StringVarP(&dbFile, "database", "d", "data.db", "database file path")
	delroomCmd.PersistentFlags().Uint64VarP(&dbId, "id", "i", 0, "table id")

	editroomCmd.PersistentFlags().StringVarP(&dbFile, "database", "d", "data.db", "database file path")
	editroomCmd.PersistentFlags().Uint64VarP(&dbId, "id", "i", 0, "table id")
	editroomCmd.PersistentFlags().StringVarP(&roomPlatform, "platform", "p", "", "Room platform")
	editroomCmd.PersistentFlags().StringVarP(&roomId, "roomid", "r", "", "Room id")
	editroomCmd.PersistentFlags().StringVarP(&roomName, "name", "n", "", "Room name")

}
