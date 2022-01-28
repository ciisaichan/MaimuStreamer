package bilibili

import "MaimuStreamer/global"

type BiliTask struct {
	Downloader global.FileDownloader
	RoomID     string
	Name       string
	Title      string
	DBId       uint64
}
