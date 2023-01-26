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

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "This command will retrieve the questions",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		response, err := http.Get("http://localhost:8080/questions/1")
		
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

			var result question
			if err := json.Unmarshal(b, &result); err != nil {  // Parse []byte to the go struct pointer
				fmt.Println("Can not unmarshal JSON")
			}
			fmt.Printf(result.Question)
        } 
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
