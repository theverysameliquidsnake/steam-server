package configs

import "log"

var logger bool

func EnableLogger() {
	logger = true
}

func DisableLogger() {
	logger = false
}

func PrintLog(message string) {
	if logger {
		log.Println(message)
	}
}
