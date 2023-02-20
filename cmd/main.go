package main

import (
	"flag"
	"votingMicroservicesApp/pkg/config"
)

func main() {
	// Parse the command line flags.
	isWorker := flag.Bool("isWorker", false, "start in worker mode")
	flag.Parse()

	// Initialize the application configuration and defer shutdown until the end of main.
	config.InitializeAppConfig(false)
	defer config.ShutDown()

	// Call the appropriate function based on the value of the isWorker flag.
	if *isWorker {
		worker()
	} else {
		web()
	}
}
