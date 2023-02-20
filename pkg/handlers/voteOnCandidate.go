package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"votingMicroservicesApp/pkg/models"
)

func VoteOnCandidate(c *gin.Context) {
	var candidateInfo CandidateInfo
	if err := c.ShouldBindUri(&candidateInfo); err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	res, err := models.FetchVotingSession(candidateInfo.SessionUuid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Voting Session Not Found"})
		return
	}
	err = res.VoteOnCandidateAction(candidateInfo.CandidateUuid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Candidate Not Found"})
		return
	}
	c.AsciiJSON(http.StatusAccepted, gin.H{"msg": "Accepted"})
}
