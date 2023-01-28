/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getQuestionsCmd = &cobra.Command{
	Use:   "start-quizz",
	Short: "Run the start-quizz command to start the quizz.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		usernamePrompt := promptContent{
			"Please enter your username",
		}
		username := promptGetInput(usernamePrompt)
		response, err := http.Get("http://localhost:8080/questions")

		if err != nil {
			fmt.Println(err)
		}

		defer response.Body.Close()

		if response.StatusCode == 200 {

			if err != nil {
				fmt.Println(err)
			}

			b, err := io.ReadAll(response.Body)
			// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
			if err != nil {
				log.Fatalln(err)
			}
			// fmt.Println(string(b))

			var quizz []question
			if err := json.Unmarshal(b, &quizz); err != nil { // Parse []byte to the go struct pointer
				fmt.Println("Can not unmarshal JSON")
			}
			answers := makeChoice(quizz)

			player := calculateScore(answers, username)

			fmt.Println(player.Name + ":" + strconv.Itoa(player.Score))

			calculateStandings(player)
		}
	},
}

// calculate the score
func calculateScore(answers []string, username string) playerStruct {

	x := 0
	for i := 0; i < len(correctAnswers); i++ {
		if correctAnswers[i].ChoiceId == answers[i] {
			x = x + 1
		}
	}
	player := playerStruct{Name: username, Score: x * 10}
	return player
}

func calculateStandings(player playerStruct) {
	response, err := http.Get("http://localhost:8080/standings")

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {

		if err != nil {
			fmt.Println(err)
		}

		b, err := io.ReadAll(response.Body)
		// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
		if err != nil {
			log.Fatalln(err)
		}
		// fmt.Println(string(b))

		var standings []playerStruct
		if err := json.Unmarshal(b, &standings); err != nil { // Parse []byte to the go struct pointer
			fmt.Println("Can not unmarshal JSON")
		}
		
		standings = append(standings, player)

		sort.SliceStable(standings, func(i, j int) bool{
   			return standings[i].Score > standings[j].Score
		})

		fmt.Println("-------------------The Results------------------")
		for i := 0; i < len(standings); i++ {
			fmt.Println(strconv.Itoa(i+1)+") "+standings[i].Name+":"+strconv.Itoa(standings[i].Score))
		} 
	}
}

func init() {
	rootCmd.AddCommand(getQuestionsCmd)
}

type promptContent struct {
	label string
}

func promptGetInput(pc promptContent) string {
	prompt := promptui.Prompt{
		Label: pc.label,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	//fmt.Printf("Input: %s\n", result)

	return result
}

func makeChoice(quizzResponse []question) []string {
	var answers = []string{}
	for _, quizz := range quizzResponse {
		question1 := promptContent{
			fmt.Sprintf(quizz.Question),
		}

		choices := []string{}

		for _, choice := range quizz.Choices {
			choices = append(choices, choice.Choice)
		}

		answer := promptGetSelect(question1, choices)

		answers = append(answers, answer)
	}

	return answers
}

func promptGetSelect(pc promptContent, choices []string) string {
	items := choices
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label: pc.label,
			Items: items,
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return strconv.Itoa(index + 1)
}
