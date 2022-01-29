package global

import (
	"os"
)

func FileSize(path string) int64 {
	s, err := os.Stat(path)
	if err != nil {
		return -1
	}
	return s.Size()
}
