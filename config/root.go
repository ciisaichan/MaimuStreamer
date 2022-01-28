package config

import (
	"MaimuStreamer/model"
	"encoding/json"
	"os"
)

var Cfg model.Config

func Init(configPath string) error {
	file, err := os.Open(configPath)
	if err != nil {
		//panic("Unable to read config file: " + err.Error())
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Cfg)
	if err != nil {
		//panic("Config file format error: " + err.Error())
		return err
	}
	return nil
}
