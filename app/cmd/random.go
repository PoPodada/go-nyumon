/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

type Data struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

var (
	images     int
	httpClient = &http.Client{Timeout: 10 * time.Second}
	apiURL     = "https://dog.ceo/api/breeds/image/random"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "A brief description of your command",

	RunE: func(cmd *cobra.Command, args []string) error {

		if images < 1 {
			images = 1
		}
		ctx := cmd.Context()

		for i := 0; i < images; i++ {
			url, err := fetchOne(ctx)
			if err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), url)
		}
		return nil
	},
}

func fetchOne(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return "", err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var data Data
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	if data.Status != "success" {
		return "", fmt.Errorf("API returned non-success status: %s", data.Status)
	}

	return data.Message, nil
}

func init() {
	rootCmd.AddCommand(randomCmd)

	randomCmd.Flags().IntVarP(&images, "images", "i", 1, "取得する画像枚数")
}
