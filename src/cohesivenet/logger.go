package cohesivenet

import (
	"fmt"
	"log"
	"strings"
)

var DEBUGLEVEL int = 10
var INFOLEVEL int = 20
var WARNLEVEL int = 30
var ERRORLEVEL int = 40

type Logger struct {
	Level int
}

// func getLevelString(level int) string {
// 	switch level {
// 	case DEBUGLEVEL:
// 		return "DEBUG"
// 	case INFOLEVEL:
// 		return "INFO"
// 	case WARNLEVEL:
// 		return "WARN"
// 	case ERRORLEVEL:
// 		return "ERROR"
// 	}
// 	return ""
// }

func getLogLevel(level string) int {
	switch levelUpper := strings.ToUpper(level); levelUpper {
	case "DEBUG":
		return DEBUGLEVEL
	case "INFO":
		return INFOLEVEL
	case "WARN":
		return WARNLEVEL
	case "ERROR":
		return ERRORLEVEL
	}
	return INFOLEVEL
}

func NewLogger(logLevel int) Logger {
	return Logger{
		Level: logLevel,
	}
}

func (logger Logger) SetLevel(logLevel int) {
	logger.Level = logLevel
}

func (logger Logger) Debug(message string) {
	if logger.Level > DEBUGLEVEL {
		return
	}

	finalMessage := fmt.Sprintf("[DEBUG] %+v", message)
	log.Println(finalMessage)
}

func (logger Logger) Info(message string) {
	if logger.Level > INFOLEVEL {
		return
	}

	finalMessage := fmt.Sprintf("[INFO] %+v", message)
	log.Println(finalMessage)
}

func (logger Logger) Warn(message string) {
	if logger.Level > WARNLEVEL {
		return
	}

	finalMessage := fmt.Sprintf("[WARN] %+v", message)
	log.Println(finalMessage)
}

func (logger Logger) Error(message string) {
	if logger.Level > ERRORLEVEL {
		return
	}

	finalMessage := fmt.Sprintf("[ERROR] %+v", message)
	log.Println(finalMessage)
}
