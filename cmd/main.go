package main

import (
	"flag"
	"votingMicroservicesApp/pkg/config"
)

func main() {
	WorkerMode := flag.Bool("isWorker", false, "start on work mode")
	config.InitializeAppConfig(false)
	defer config.ShutDown()
	flag.Parse()
	if *WorkerMode {
		worker()
	} else {
		web()
	}
}
