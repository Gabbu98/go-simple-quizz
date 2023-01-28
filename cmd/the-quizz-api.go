package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getQuestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database)
}

func getStandings(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, standings)
}