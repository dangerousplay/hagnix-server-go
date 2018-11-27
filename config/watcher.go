package config

import "github.com/fsnotify/fsnotify"

var watcher *fsnotify.Watcher

func initWatcher() {
	watchs, err := fsnotify.NewWatcher()

	if err != nil {
		panic(err)
	}

	watchs.Add("/app")

	watcher = watchs

	go watch()
}

func watch() {
	for {
		event, ok := <-watcher.Events

		if event.Name == "" && ok {

		}
	}
}
