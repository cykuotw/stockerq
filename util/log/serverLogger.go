package log

import (
	"log/slog"
	"os"
	"time"
)

var serverFileLogger *fileLogger
var serverTermLogger *terminalLogger

func LogServer(event string, api string, method string, body map[string]string) {
	fileEnable := os.Getenv("logServerTermEn") == "True"
	if fileEnable {
		logServerFile(event, api, method, body)
	}

	termEnable := os.Getenv("logServerFileEn") == "True"
	if termEnable {
		logServerTerm(event, api, method, body)
	}
}

func logServerFile(event string, api string, method string, body map[string]string) {
	filename := "server-info-" + time.Now().Format("2006-01-02") + ".log"

	// create new file if logger is nil
	//              or if filename is not expected
	if serverFileLogger == nil || serverFileLogger.Filename != filename {
		file, err := newLoggerFile(filename)
		if err != nil {
			// TODO
			return
		}

		logger := slog.New(slog.NewJSONHandler(file, nil))
		serverFileLogger = &fileLogger{
			Filename: filename,
			Logger:   logger,
		}

		serverFileLogger.Logger.WithGroup("detail").Info(event,
			"api", api,
			"method", method,
			"body", body,
		)
	}
}

func logServerTerm(event string, api string, method string, body map[string]string) {
	if serverTermLogger == nil {
		serverTermLogger = &terminalLogger{
			Logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
		}
	}

	serverTermLogger.Logger.WithGroup("detail").Info(event,
		"api", api,
		"method", method,
		"body", body,
	)
}
