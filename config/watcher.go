package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/ivahaev/go-logger"
	"io/ioutil"
	"os"
	"strings"
)

var watcher *fsnotify.Watcher

func initWatcher() {
	watchs, err := fsnotify.NewWatcher()

	if err != nil {
		panic(err)
	}

	watchs.Add(executablePath + "/app")
	watchs.Add(executablePath + "/app/Languages")

	watcher = watchs

	reload()

	go watch()
}

func watch() {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				break
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				logger.Info("Changes detected on config, reloading...")
				reload()
				logger.Info("Reload completed!")
			}
			break
		case error, ok := <-watcher.Errors:
			if !ok {
				break
			}
			logger.Warn("Event watcher error: " + error.Error())
			break
		}
	}
}

func reload() {
	init, err := os.Open(executablePath + "/app/init.txt")

	if err != nil {
		logger.Info("Can't load /app/init.txt: " + err.Error())
	} else {
		bytes, err := ioutil.ReadAll(init)
		if err != nil {
			logger.Info("Can't load /app/init.txt: " + err.Error())
		} else {
			filesConfig.Init = string(bytes)
		}
	}

	info, err := ioutil.ReadDir(executablePath + "/app/Languages")

	if filesConfig.Languages == nil {
		filesConfig.Languages = map[string]string{}
	}

	if err != nil {
		logger.Info("can't load /app/Languages: " + err.Error())
	} else {
		for _, v := range info {
			if !v.IsDir() {
				file, err := os.Open(executablePath + "/app/Languages/" + v.Name())

				if err != nil {
					logger.Warn("Can't open file: " + v.Name())
					continue
				}

				lines, err := ioutil.ReadAll(file)

				if err != nil {
					logger.Warn("Can't load file: " + err.Error())
					continue
				}

				index := strings.Replace(v.Name(), ".txt", "", 1)

				filesConfig.Languages[index] = string(lines)
			}
		}
	}

}
