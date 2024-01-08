package log

import (
	"log/slog"
	"os"
	apperror "stockerq/web/app/app-error"
	"time"
)

var errFileLogger *fileLogger
var errTermLogger *terminalLogger

func LogError(layer string, err *apperror.AppError) {
	fileEnable := os.Getenv("logErrorFileEn") == "True"
	if fileEnable {
		logErrorFile(layer, err)
	}

	termEnable := os.Getenv("logErrorTermEn") == "True"
	if termEnable {
		logErrorTerm(layer, err)
	}
}

func logErrorFile(layer string, err *apperror.AppError) {
	filename := "err-" + time.Now().Format("2006-01-02") + ".log"

	// create new file if logger is nil
	//              or if filename is not expected
	if errFileLogger == nil || errFileLogger.Filename != filename {
		file, err := newLoggerFile(filename)
		if err != nil {
			// TODO
			return
		}

		logger := slog.New(slog.NewJSONHandler(file, nil))
		errFileLogger = &fileLogger{
			Filename: filename,
			Logger:   logger,
		}
	}
	errFileLogger.Logger.WithGroup("detail").Error(err.Err.Error(),
		"layer", layer,
		"file", err.CallerFile,
		"line", err.CallerLine,
		"function", err.CallerFunction,
	)
}

func logErrorTerm(layer string, err *apperror.AppError) {
	if errTermLogger == nil {
		errTermLogger = &terminalLogger{
			Logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
		}
	}

	errTermLogger.Logger.Error(err.Error())
}
