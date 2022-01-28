package model

type LiveRoom struct {
	ID       uint64 `gorm:"primaryKey;column:id;index" json:"id"`
	Platform string `gorm:"not null;column:platform" json:"platform"`
	RoomID   string `gorm:"not null;column:roomid" json:"roomid"`
	Name     string `gorm:"not null;column:name" json:"name"`
	TimeStamp
}
