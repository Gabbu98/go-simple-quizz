package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getQuestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database)
}

// postAnswers
func postAnswer(c *gin.Context) {
	var answer answer

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&answer); err != nil {
		return
	}

	// Add the answer to the slice.
	answers = append(answers, answer)
	c.IndentedJSON(http.StatusCreated, answer)

}