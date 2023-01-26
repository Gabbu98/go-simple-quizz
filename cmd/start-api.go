/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// startQuizzCmd represents the startQuizz command
var startQuizzCmd = &cobra.Command{
	Use:   "start-api",
	Short: "Run the start-api command to run the api for the quizz. Must be executed first!",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Quizz Api Starting called")
		var router = gin.Default()
		router.GET("/questions", getQuestions)
		router.POST("/submit", postAnswer)
		router.Run("localhost:8080")
	},
}

func init() {
	rootCmd.AddCommand(startQuizzCmd)
}
