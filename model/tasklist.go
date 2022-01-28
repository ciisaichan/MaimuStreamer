package model

type TaskList struct {
	ID          uint64 `gorm:"primaryKey;column:id;index" json:"id"`
	Platform    string `gorm:"not null;column:platform" json:"platform"`
	RoomID      string `gorm:"not null;column:roomid" json:"roomid"`
	Name        string `gorm:"not null;column:name" json:"name"`
	RoomTitle   string `gorm:"not null;column:room_title" json:"room_title"`
	DLByteCount int64  `gorm:"not null;column:dl_byte_count" json:"dl_byte_count"`
	TimeStamp
}
