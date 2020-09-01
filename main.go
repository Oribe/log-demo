package main

import "Demo/exercise/logger/logger"

func main() {
	log := logger.NewLogger("debug", "")

	for {

		// log.Debug("这是一条Debug信息")
		// log.Info("这是一条Info信息")
		// log.Warning("这是一条Warning信息")
		log.Error("这是一条Error信息")
		// log.Fatal("这是一条Fatal信息")
	}
}
