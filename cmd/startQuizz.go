/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// startQuizzCmd represents the startQuizz command
var startQuizzCmd = &cobra.Command{
	Use:   "startQuizz",
	Short: "The Quizz Api",
	Long: `Run the startQuizz command to run the api for the quizz.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Quizz Api Starting called")
	},
}

func init() {
	rootCmd.AddCommand(startQuizzCmd)
	var router = gin.Default()
	router.GET("/questions", getQuestions)
	router.POST("/questions/:id", postAnswer)
	router.Run("localhost:8080")
	
}

type question struct {
	Id       int      `json:"id"`
	Question string   `json:"question"`
	Choices  []Choice `json:"choices"`
}

type option struct {
	QuestionId int `json:"questionId"`
}

type Choice struct {
	Id     int    `json:"id"`
	Choice string `json:"choice"`
}

type correctAnswer struct {
	QuestionId int `json:"questionId"`
	ChoiceId   int `json:"choiceId"`
}

type answer struct {
	QuestionId int `json:"questionId"`
	ChoiceId   int `json:"choiceId"`
}

var questions = []question{
	{Id: 1,
		Question: "What country has the highest life expectancy?", Choices: []Choice{
			{Id: 1, Choice: "Spain"},
			{Id: 2, Choice: "Hong Kong"},
			{Id: 3, Choice: "Italy"},
			{Id: 4, Choice: "Malta"},
		}},
}

var options = []option{
	{QuestionId: 1}}

var correctAnswers = []answer{
	{QuestionId: 1, ChoiceId: 2},
}

var answers = []answer{
	
}

func getQuestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, questions)
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
