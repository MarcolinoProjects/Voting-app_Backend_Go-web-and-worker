package main

import (
	"log"
	"votingMicroservicesApp/pkg/config"
	"votingMicroservicesApp/pkg/models"
)

func worker() {
	forever := make(chan bool)
	msgs := config.AppContext.RabbitConfig.ListenToQueue("voting")
	go func() {
		for d := range msgs {
			votingAction, err := models.UnmarshalVotingAction(d.Body)
			if err != nil {
				log.Panic(err)
				return
			}
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
