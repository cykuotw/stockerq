package util

import (
	"runtime"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func HandleError(err error, errorMsg string) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Warn(errorMsg +
			"\n\t source: " + file +
			" (line: " + strconv.Itoa(line) + ")" +
			"\n\tmessage: " + err.Error())
	}
}
