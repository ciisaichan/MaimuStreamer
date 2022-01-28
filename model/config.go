package model

type Config struct {
	SqliteDb            string `json:"sqlite_db"`
	StorageDir          string `json:"storage_dir"`
	RoomsCheckDelay     int    `json:"rooms_check_delay"`
	BiliTasksCheckDelay int    `json:"bili_tasks_check_delay"`
}