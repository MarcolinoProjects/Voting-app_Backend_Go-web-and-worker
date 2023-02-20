package main

import (
	"log"
	"votingMicroservicesApp/pkg/config"
	"votingMicroservicesApp/pkg/models"
)

// worker is the entry point for the worker mode of the application.
func worker() {
	// Create a channel to wait for the worker to finish.
	forever := make(chan bool)

	// Listen to the voting queue for messages.
	msgs := config.AppContext.RabbitConfig.ListenToQueue("voting")

	// Process messages received from the voting queue.
	go func() {
		for d := range msgs {
			// Unmarshal the voting action from the message body.
			votingAction, err := models.UnmarshalVotingAction(d.Body)
			if err != nil {
				log.Panic(err)
				return
			}

			// Count the vote based on the voting action.
			err = models.CountVote(votingAction)
			if err != nil {
				log.Panic(err)
				return
			}

			log.Printf("Received a message: %s", votingAction)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
