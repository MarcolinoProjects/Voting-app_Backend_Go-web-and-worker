package handlers

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {

	route := gin.Default()
	route.Use(CORSMiddleware)
	api := route.Group("/api")
	{
		api.POST("/create", CreateVoting)
		api.GET("/list", ListVoting)
		api.GET("/:sessionUuid", GetVotingSessionInfo)
		api.DELETE("/:sessionUuid", DeleteVotingSession)
		api.POST("/:sessionUuid/:candidateUuid", VoteOnCandidate)
	}

	return route
}
func CORSMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, ResponseType, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
}
