package helpers

import (
	"time"
)

func GetTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetDate() string {
	return time.Now().Format("2006-01-02")
}


