package handlers

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/nu7hatch/gouuid"
	"log"
	"net/http"
	"votingMicroservicesApp/pkg/models"
)

/*
CreateVoting this is the handler for create new models.Voting sessions
example input:

POST /... HTTP/1.1
Host: ...
Content-Type: application/json
Content-Length: 368

	{
	  "uuid": "ba9ef00d-3cd6-48a5-8f55-1d0fd00eebd6",
	  "name": "voting",
	  "candidates": [

	      {
	        "uuid": "36731832-72a6-4811-a61b-5b20cfd5bfd2",
	        "name": "Roberts",
	        "votes": 0
	      },
	      {
	        "uuid": "e7b35078-8eec-4d74-ad14-d5654d95c773",
	        "name": "Park",
	        "votes": 0
	      }
	  ]
	}
*/
func CreateVoting(c *gin.Context) {
	var voting models.Voting
	err := c.ShouldBind(&voting)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		log.Println(err)
	} else {
		u, _ := uuid.NewV4()
		voting.UUID = u.String()
		for idx, _ := range voting.Candidates {
			u, _ := uuid.NewV4()
			voting.Candidates[idx].UUID = u.String()
			voting.Candidates[idx].Votes = 0
		}
		voting.PersistOnRedis()
		c.AsciiJSON(http.StatusCreated, voting)
	}
}
