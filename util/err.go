package util

import (
	"runtime"
	"strconv"

	log "github.com/sirupsen/logrus"
	"go.deanishe.net/env"
)

func HandleError(err error, errorMsg string) {
	enableDisplay := env.GetBool("errorMsgOn")
	if err != nil && enableDisplay {
		_, file, line, _ := runtime.Caller(1)
		log.Warn(errorMsg +
			"\n\t source: " + file +
			" (line: " + strconv.Itoa(line) + ")" +
			"\n\tmessage: " + err.Error())
	}
}
