package config

import (
	"fmt"
	"os"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

var (
	Logger      = logrus.New()
	currentDate string
	logFile     *os.File
)

func InitLogger() {
	currentDate = time.Now().Format("2006-01-02")
	setLogFile(currentDate)

	c := cron.New()
	c.AddFunc("0 0 * * *", func() {
		newDate := time.Now().Format("2006-01-02")
		if newDate != currentDate {
			currentDate = newDate
			setLogFile(currentDate)
		}
	})
	c.Start()
}

func setLogFile(date string) {
	if logFile != nil {
		logFile.Close()
	}

	os.MkdirAll("logs", os.ModePerm)

	fileName := fmt.Sprintf("logs/%s.log", date)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal("Failed to open log file: ", err)
	}

	Logger.SetOutput(file)
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05",
		ForceColors:      false,
		DisableColors:    true,
		QuoteEmptyFields: true,
	})
	Logger.SetLevel(logrus.InfoLevel)

	logFile = file
}

func LogInfo(msg string, fields map[string]interface{}) {
	Logger.WithFields(fields).Info(msg)
}

func LogError(msg string, fields map[string]interface{}) {
	Logger.WithFields(fields).Error(msg)
}

func LogWarning(msg string, fields map[string]interface{}) {
	Logger.WithFields(fields).Warn(msg)
}

func LogCritical(msg string, fields map[string]interface{}) {
	Logger.WithFields(fields).Fatal(msg)
}
