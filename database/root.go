package database

import (
	"MaimuStreamer/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var MainDB *gorm.DB

func Init(dbpath string) error {
	loggerSetting := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:              time.Second,
			LogLevel:                   logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	db, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger: loggerSetting,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return err
	}

	if err = db.AutoMigrate(&model.LiveRoom{}, &model.LiveRoom{}); err != nil {
		return err
	}
	if err = db.AutoMigrate(&model.TaskList{}, &model.TaskList{}); err != nil {
		return err
	}

	MainDB = db
	return nil
}