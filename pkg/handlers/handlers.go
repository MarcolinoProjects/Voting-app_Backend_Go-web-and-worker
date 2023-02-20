package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"votingMicroservicesApp/pkg/models"
)

func ListVoting(c *gin.Context) {
	models.FetchVotingSessions()
	c.String(http.StatusOK, "OKAY")

}
