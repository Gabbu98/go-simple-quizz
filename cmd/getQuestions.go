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
	Use:   "getQuestions",
	Short: "This command will retrieve the questions",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		createNewNote()
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
			//fmt.Printf(result[0].Question)
        } 
	},
}

func init() {
	rootCmd.AddCommand(getQuestionsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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

func createNewNote() {
    // wordPromptContent := promptContent{
    //     "Please provide a number.",
    //     "What country has the highest life expectancy?",
    // }
    // word := promptGetInput(wordPromptContent)

	question1 := promptContent{
        "Please Try again",
        fmt.Sprintf("With respect to the certification of airmen, which is a category of aircraft?"),
    }
    answer1 := promptGetSelect(question1, []string{"A. Gyroplane, helicopter, airship, free balloon.", "B. Airplane, rotorcraft, glider, lighter-than-air.", "C. Single-engine land and sea, multiengine land and sea."})

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