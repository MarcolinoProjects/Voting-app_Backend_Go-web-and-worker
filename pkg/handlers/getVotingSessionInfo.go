package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"votingMicroservicesApp/pkg/models"
)

func GetVotingSessionInfo(c *gin.Context) {
	var candidateInfo SessionInfo
	if err := c.ShouldBindUri(&candidateInfo); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	res, err := models.FetchVotingSession(candidateInfo.SessionUuid)
	if err == nil {
		c.AsciiJSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Voting Session Not Found"})
	}
}

func DeleteVotingSession(c *gin.Context) {
	var candidateInfo SessionInfo
	if err := c.ShouldBindUri(&candidateInfo); err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	res, err := models.FetchVotingSession(candidateInfo.SessionUuid)
	if err == nil {
		res.DeleteOnRedis()
		c.AsciiJSON(http.StatusOK, gin.H{"msg": "DELETED"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Voting Session Not Found"})
	}
}
