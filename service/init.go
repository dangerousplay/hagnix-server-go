package service

import (
	"github.com/InVisionApp/go-logger"
	"github.com/jasonlvhit/gocron"
)

var logger = log.NewSimple()
var scheduler = gocron.NewScheduler()

func Init() {
	logger.Info("Starting services...")
	startTasks()
}

func startTasks() {
	scheduler.Every(2).Seconds().Do(updateServers)
	scheduler.Start()
}
