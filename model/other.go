package model

type TimeStamp struct {
	Created int64 `gorm:"not null;autoCreateTime:milli;column:created_time" json:"created_time"`
	Updated int64 `gorm:"not null;autoUpdateTime:milli;column:updated_time" json:"updated_time"`
}