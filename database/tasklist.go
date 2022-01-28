package database

import (
	"MaimuStreamer/model"
)

func GetTask(id uint64) (*model.TaskList, error) {
	task := model.TaskList{}
	if err := MainDB.First(&task, id).Limit(1).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func AddTask(platform string, roomid string, name string, roomtitle string, dlbytecount int64) (*model.TaskList, error) {
	task := model.TaskList{
		RoomID:      roomid,
		Platform:    platform,
		Name:        name,
		RoomTitle:   roomtitle,
		DLByteCount: dlbytecount,
	}
	if err := MainDB.Create(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func UpdateTask(id uint64, platform string, roomid string, name, roomtitle string, dlbytecount int64) (*model.TaskList, error) {
	task := model.TaskList{ID: id}
	if err := MainDB.First(&task).Limit(1).Error; err != nil {
		return nil, err
	}
	if err := MainDB.Model(&task).Updates(map[string]interface{}{"platform": platform, "roomid": roomid, "name": name, "room_title": roomtitle, "dl_byte_count": dlbytecount}).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func DelTask(id uint64) error {
	if err := MainDB.First(&model.TaskList{ID: id}).Limit(1).Error; err != nil {
		return err
	}
	if err := MainDB.Where("id = ?", id).Delete(&model.TaskList{}).Error; err != nil {
		return err
	}
	return nil
}

func GetAllTasks() (*[]model.TaskList, error) {
	var tasks []model.TaskList
	if err := MainDB.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return &tasks, nil
}
