/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getQuestionsCmd = &cobra.Command{
	Use:   "start-quizz",
	Short: "Run the start-quizz command to start the quizz.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		
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

			var result []question
			if err := json.Unmarshal(b, &result); err != nil {  // Parse []byte to the go struct pointer
				fmt.Println("Can not unmarshal JSON")
			}
			createNewNote(result)
			//fmt.Printf(result[0].Question)
        } 
	},
}

func init() {
	rootCmd.AddCommand(getQuestionsCmd)
}

type promptContent struct {
    errorMsg string
    label    string
}

func promptGetInput(pc promptContent) string {
    validate := func(input string) error {
        if len(input) <= 0 {
            return errors.New(pc.errorMsg)
        }
        return nil
    }

    templates := &promptui.PromptTemplates{
        Prompt:  "{{ . }} ",
        Valid:   "{{ . | green }} ",
        Invalid: "{{ . | red }} ",
        Success: "{{ . | bold }} ",
    }

    prompt := promptui.Prompt{
        Label:     pc.label,
        Templates: templates,
        Validate:  validate,
    }

    result, err := prompt.Run()
    if err != nil {
        fmt.Printf("Prompt failed %v\n", err)
        os.Exit(1)
    }

    fmt.Printf("Input: %s\n", result)

    return result
}

func createNewNote(result []question) {
    // wordPromptContent := promptContent{
    //     "Please provide a number.",
    //     "What country has the highest life expectancy?",
    // }
    // word := promptGetInput(wordPromptContent)

	question1 := promptContent{
        "Please Try again",
        fmt.Sprintf(result[0].Question),
    }
    answer1 := promptGetSelect(question1, []string{result[0].Choices[0].Choice,result[0].Choices[1].Choice,result[0].Choices[2].Choice})

	fmt.Printf(answer1)

	// question2 := promptContent{
	// 	"Please Try again",
    //     fmt.Sprintf("1. What country has the highest life expectancy?"),
	// }
}

func promptGetSelect(pc promptContent, choices []string) string {
    items := choices
    index := -1
    var result string
    var err error

    for index < 0 {
        prompt := promptui.SelectWithAdd{
            Label:    pc.label,
            Items:    items,
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

    fmt.Printf("Input: %s\n", strconv.Itoa(index))

    return result
}