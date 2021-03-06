package main

import (
	"github.com/charllossdev/kafka-go-gin/app"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"Application": "TestApp",
		"Version:":    "v0.1.1",
	}).Info("Application start...")
	myapp := app.New("test app")
	myapp.Run()
}
