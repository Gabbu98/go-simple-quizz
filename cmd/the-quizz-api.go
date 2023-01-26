package cmd

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getQuestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database)
}

// postAnswers
func postAnswer(c *gin.Context) {
	var answers []answer

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&answers); err != nil {
		return
	}

	// Add the answer to the slice.
	//answers = append(answers, answers)
	c.IndentedJSON(http.StatusCreated, answers)
	x := 0
	for i := 0; i < len(correctAnswers); i++ {
		if correctAnswers[i].ChoiceId == answers[i].ChoiceId {
			x = x + 1
		}
	}
	fmt.Printf("Correct answers are %s", strconv.Itoa(x))
}