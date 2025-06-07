package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

// var (
// 	InfoLog  *log.Logger
// 	ErrorLog *log.Logger
// )

// func InitLoggers() {
// 	green := "\033[32m"
// 	red := "\033[31m"
// 	reset := "\033[m"

// 	InfoLog = log.New(os.Stdout, green+"INFO\t"+reset, log.Ldate|log.Ltime)
// 	ErrorLog = log.New(os.Stdout, red+"ERROR\t"+reset, log.Ldate|log.Ltime|log.Lshortfile)

// }

var Logger = logrus.New()

func InitLoggers(format string) {
	Logger.SetOutput(os.Stdout)

	if format == "text" {
		Logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	} else {
		Logger.SetFormatter(&logrus.JSONFormatter{})
	}

}
