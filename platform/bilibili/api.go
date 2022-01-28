package bilibili

import (
	"MaimuStreamer/global"
	"MaimuStreamer/model/bilibili"
	"encoding/json"
	"strconv"
)

const BILI_API_URL string = "https://api.live.bilibili.com"

func GetRoomInfo(rid string) (bilibili.RoomInfo, error) {
	var roomInfo bilibili.RoomInfo
	respByte, err := global.HttpGet(BILI_API_URL+"/room/v1/Room/get_info?id="+rid, nil)
	if err != nil {
		return roomInfo, err
	}

	err = json.Unmarshal(respByte, &roomInfo)
	if err != nil {
		return roomInfo, err
	}

	return roomInfo, nil
}

func GetPlayURL(cid int) (bilibili.PlayUrl, error) {
	var playUrl bilibili.PlayUrl
	respByte, err := global.HttpGet(BILI_API_URL+"/room/v1/Room/playUrl?cid="+strconv.Itoa(cid)+"&quality=4&platform=web", nil)

	if err != nil {
		return playUrl, err
	}

	err = json.Unmarshal(respByte, &playUrl)
	if err != nil {
		return playUrl, err
	}

	return playUrl, nil
}
