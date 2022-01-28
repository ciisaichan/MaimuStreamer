package logger

import (
	go_logger "github.com/phachon/go-logger"
)

var L *go_logger.Logger

func init()  {
	L = go_logger.NewLogger()

	L.Detach("console")

	// console adapter config
	consoleConfig := &go_logger.ConsoleConfig{
		Color: true,
		JsonFormat: false,
		Format: "%millisecond_format% [%level_string%] [%file%:%line%] %body%",
	}
	// add output to the console
	L.Attach("console", go_logger.LOGGER_LEVEL_DEBUG, consoleConfig)

}
