/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

type Data struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		var data Data
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			log.Fatal(err)
		}

		if data.Status != "success" {
			log.Fatalf("API returned non-success status: %s", data.Status)
		}

		fmt.Println(data.Message)
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
