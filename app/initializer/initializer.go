package initializer

import (
	"ic-indexer-worker/app/config"
	"log"
)

func Initialize() {

	initializeConfig()

}

func initializeConfig() {

	err := config.Initialize()
	config.InitializeKafkaConsumer()

	if err != nil {
		log.Fatal(nil, "error initializing Config :", err)
		return
	}

}
