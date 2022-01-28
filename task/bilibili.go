package task

import (
	"MaimuStreamer/config"
	"MaimuStreamer/database"
	"MaimuStreamer/logger"
	"MaimuStreamer/model"
	bili_model "MaimuStreamer/model/bilibili"
	"MaimuStreamer/platform/bilibili"
	"path"
	"time"
)

var biliTaskList map[string]*bili_model.BiliTask

func BiliRoomCheck(room model.LiveRoom) {
	if _, exist := biliTaskList[room.RoomID]; exist {
		return
	}
	roomInfo, err := bilibili.GetRoomInfo(room.RoomID)
	if err != nil {
		logger.L.Error("[BiliBili] Unable to request room info: " + err.Error())
	}

	if roomInfo.Code == 0 {
		if roomInfo.Data.LiveStatus == 1 {
			playUrl, err := bilibili.GetPlayURL(roomInfo.Data.RoomID)
			if err != nil {
				logger.L.Error("[BiliBili] Unable to request playurl: " + err.Error())
			}
			if playUrl.Code == 0 {
				biliTaskList[room.RoomID] = new(bili_model.BiliTask)
				biliTaskList[room.RoomID].RoomID = room.RoomID
				biliTaskList[room.RoomID].Name = room.Name
				biliTaskList[room.RoomID].Title = roomInfo.Data.Title
				biliTaskList[room.RoomID].Downloader.Url = playUrl.Data.Durl[0].URL
				biliTaskList[room.RoomID].Downloader.FileName = path.Join(config.Cfg.StorageDir, "bilibili", room.Name+" "+time.Now().Format("2006-01-02 15-04-05")+".flv")
				biliTaskList[room.RoomID].Downloader.Start()
			} else {
				logger.L.Error("[BiliBili] Unable to get playurl: " + playUrl.Message)
			}
		} else {
			//logger.L.Info("[BiliBili] " + room.Name + ": Not live.")
		}
	} else {
		logger.L.Error("[BiliBili] Unable to get room info: " + roomInfo.Message)
	}
}

func BiliTaskCheckLoop() {
	for {
		BiliTaskListCheck()
		time.Sleep(time.Duration(config.Cfg.BiliTasksCheckDelay) * time.Millisecond)
	}
}

func BiliTaskListCheck() {
	for k, v := range biliTaskList {
		if v.DBId == 0 {
			if v.Downloader.Downloading {
				if v.Downloader.HttpResponse != nil {
					task, err := database.AddTask("bilibili", v.RoomID, v.Name, v.Title, v.Downloader.Reader.Current)
					if err != nil {
						logger.L.Error("[BiliBili] Unable add task to database: " + err.Error())
					} else {
						biliTaskList[k].DBId = task.ID
						poolWait.Add(1)
						logger.L.Info("[BiliBili] " + v.Name + ": Start Record.")
					}
				}
			} else {
				if v.Downloader.Error != nil {
					logger.L.Error("[BiliBili] Unknown error in Downloader: " + v.Downloader.Error.Error())
					delete(biliTaskList, k)
					poolWait.Done()
				}
			}
		} else {
			if !v.Downloader.Downloading {
				if v.Downloader.Error != nil {
					logger.L.Error("[BiliBili] Task " + v.Name + " error occurred: " + v.Downloader.Error.Error())
				} else {
					logger.L.Info("[BiliBili] " + v.Name + ": Record complete.")
				}
				err := database.DelTask(v.DBId)
				if err != nil {
					logger.L.Error("[BiliBili] Unable to delete task " + v.Name + " from database: " + err.Error())
				}
				delete(biliTaskList, k)
				poolWait.Done()
			} else {
				if _, exist := database.ExistRoom(v.RoomID); !exist {
					v.Downloader.Stop()
				}
				_, err := database.UpdateTask(v.DBId, "bilibili", v.RoomID, v.Name, v.Title, v.Downloader.Reader.Current)
				if err != nil {
					logger.L.Error("[BiliBili] Unable to update task " + v.Name + " from database: " + err.Error())
				}
			}
		}
	}
}

func biliTaskExit() {
	for _, v := range biliTaskList {
		if v.Downloader.HttpResponse != nil{
			v.Downloader.Stop()
		}
	}
}
