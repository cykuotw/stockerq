package log

import (
	"log/slog"
	"os"
)

type fileLogger struct {
	Filename string
	Logger   *slog.Logger
}

type terminalLogger struct {
	Logger *slog.Logger
}

func newLoggerFile(filename string) (*os.File, error) {
	dir := os.Getenv("logDir")
	if dir == "" {
		dir = os.Getenv("HOME") + "/log/stockerq/"
	}

	dirExist := false
	_, err := os.Stat(dir)
	if err == nil {
		dirExist = true
	}
	if os.IsNotExist(err) {
		dirExist = false
	}

	if !dirExist {
		err := os.MkdirAll(dir, 0770)
		if err != nil {
			return nil, err
		}
	}

	file, err := os.OpenFile(
		dir+"/"+filename,
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		0660)
	if err != nil {
		return nil, err
	}

	return file, nil
}
