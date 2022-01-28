package database

import (
	"MaimuStreamer/model"
	"errors"
	"strconv"
)

func GetLiveRoom(id uint64) (*model.LiveRoom, error) {
	room := model.LiveRoom{}
	if err := MainDB.First(&room, id).Limit(1).Error; err != nil {
		return nil, err
	}
	return &room, nil
}

func AddLiveRoom(platform string, roomid string, name string) (*model.LiveRoom, error) {
	id, exist := ExistRoom(roomid)
	if exist {
		return nil, errors.New("Room already exists, ID: " + strconv.FormatUint(id, 10))
	}
	room := model.LiveRoom{
		RoomID:   roomid,
		Platform: platform,
		Name:     name,
	}
	if err := MainDB.Create(&room).Error; err != nil {
		return nil, err
	}
	return &room, nil
}

func UpdateLiveRoom(id uint64, platform string, roomid string, name string) (*model.LiveRoom, error) {
	r := model.LiveRoom{ID: id}
	if err := MainDB.First(&r).Limit(1).Error; err != nil {
		return nil, err
	}
	if err := MainDB.Model(&r).Updates(map[string]interface{}{"platform": platform, "roomid": roomid, "name": name}).Error; err != nil {
		return nil, err
	}
	return &r, nil
}

func DelLiveRoom(id uint64) error {
	if err := MainDB.First(&model.LiveRoom{ID: id}).Limit(1).Error; err != nil {
		return err
	}
	if err := MainDB.Where("id = ?", id).Delete(&model.LiveRoom{}).Error; err != nil {
		return err
	}
	return nil
}

func GetAllRooms() (*[]model.LiveRoom, error) {
	var rooms []model.LiveRoom
	if err := MainDB.Find(&rooms).Error; err != nil {
		return nil, err
	}
	return &rooms, nil
}

func ExistRoom(roomid string) (uint64, bool) {
	r := model.LiveRoom{}
	if err := MainDB.First(&r, "roomid = ?", roomid).Limit(1).Error; err != nil {
		return 0, false
	}
	return r.ID, true
}
