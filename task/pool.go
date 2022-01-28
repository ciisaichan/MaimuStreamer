package task

import (
	"MaimuStreamer/config"
	"MaimuStreamer/database"
	"MaimuStreamer/global"
	"MaimuStreamer/logger"
	bili_model "MaimuStreamer/model/bilibili"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

var poolWait *sync.WaitGroup

func init() {
	biliTaskList = make(map[string]*bili_model.BiliTask)
}

func Init() {
	CheckTasksTable()
	poolWait = new(sync.WaitGroup)
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				if !global.Aborting {
					ExitPool()
				}
			}
		}
	}()

	logger.L.Info("[Task Pool] Start running...")
	go BiliTaskCheckLoop()
	go TaskLoop()

	poolWait.Add(1)
	poolWait.Wait()
	logger.L.Info("[Task Pool] Goodbye ~")

}

func TaskLoop() {
	for !global.Aborting {
		logger.L.Info("[Task Pool] Checking live rooms...")
		RoomsCheck()
		time.Sleep(time.Duration(config.Cfg.RoomsCheckDelay) * time.Millisecond)
	}
}

func RoomsCheck() {
	rooms, err := database.GetAllRooms()
	if err != nil {
		logger.L.Error("[Task Pool] Unable to read all rooms: " + err.Error())
		return
	}
	for _, room := range *rooms {
		switch room.Platform {
		case "bilibili":
			go BiliRoomCheck(room)
		case "youtube":
			//logger.L.Info("test youtube")
		}
	}
}

func ExitPool() {
	global.Aborting = true
	logger.L.Info("[Task Pool] Exiting task pool...")
	poolWait.Done()

	go biliTaskExit()
}

func CheckTasksTable() {
	tasks, err := database.GetAllTasks()
	if err != nil {
		fmt.Println("Unable to read all tasks: " + err.Error())
	}
	var input string
	if len(*tasks) > 0 {
		fmt.Print("Some residual data in the task list. Do you want to clear it? (Y/N): ")
		fmt.Scan(&input)
		if strings.EqualFold(input, "y") {
			for _, v := range *tasks {
				err := database.DelTask(v.ID)
				if err != nil {
					fmt.Println("Unable to delete task: " + err.Error())
				}
			}
			fmt.Println("Done.")
		}else {
			fmt.Println("Ignore.")
		}
		fmt.Print("\n")
	}
}
