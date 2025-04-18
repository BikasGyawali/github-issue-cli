package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

const githubAPI = "https://api.github.com/repos/%s/%s/issues"

type ListIssue struct {
	Title  string `json:"title"`
	Number int    `json:"number"`
	State  string `json:"state"`
}

type Issue struct {
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Assignees []string `json:"assignees,omitempty"`
	Labels    []string `json:"labels,omitempty"`
	Milestone int      `json:"milestone,omitempty"`
}

var owner, repo, title, body string
var assignees, labels []string
var milestone int

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Fetches the list of issues",
	Run: func(cmd *cobra.Command, args []string) {
		url := fmt.Sprintf(githubAPI, owner, repo)

		resp, err := http.Get(url)

		if err != nil {
			log.Fatalf("Error fetching response: %v", err)
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Fatalf("GitHub API returned error: %s", resp.Status)
		}

		var issues []ListIssue
		if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
			log.Fatalf("Error decoding response: %v", err)
		}

		// Display issues
		if len(issues) == 0 {
			fmt.Println("No issues found.")
			return
		}

		for _, issue := range issues {
			fmt.Printf("#%d - %s [%s]\n", issue.Number, issue.Title, issue.State)
		}

	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an issue",
	Run: func(cmd *cobra.Command, args []string) {
		if body == "" {
			prompt := promptui.Prompt{
				Label: "Enter issue body",
			}
			userInput, err := prompt.Run()
			if err != nil {
				log.Fatalf("Prompt failed: %v", err)
			}
			body = userInput
		}

		var issueDetails = Issue{
			Title:     title,
			Body:      body,
			Assignees: assignees,
			Labels:    labels,
			Milestone: milestone,
		}

		jsonData, err := json.Marshal(issueDetails)

		if err != nil {
			log.Fatalf("Error converting to json: %v", err)
		}

		url := fmt.Sprintf(githubAPI, owner, repo)

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))

		if err != nil {
			log.Fatalf("Error creating request: %v", err)
		}

		authorization := "token " + os.Getenv("GITHUB_PAT")

		req.Header.Set("Authorization", authorization)
		req.Header.Set("Accept", "application/vnd.github.v3+json")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("Error sending request: %v", err)
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			log.Fatalf("Failed to create issue: %s", resp.Status)
		}

		fmt.Println("Issue created successfully!")

	},
}

func init() {
	rootCmd.AddCommand(listCmd, createCmd)
	//flags for listing issues
	listCmd.Flags().StringVarP(&owner, "owner", "o", "", "GitHub repo owner")
	listCmd.Flags().StringVarP(&repo, "repo", "r", "", "GitHub repo name")

	//flags for creating issues
	createCmd.Flags().StringVarP(&owner, "owner", "o", "", "GitHub repo owner")
	createCmd.Flags().StringVarP(&repo, "repo", "r", "", "GitHub repo name")
	createCmd.Flags().StringVarP(&title, "title", "t", "", "Issue title (required)")
	createCmd.Flags().StringVarP(&body, "body", "b", "", "Issue body description")
	createCmd.Flags().StringSliceVarP(&assignees, "assignees", "a", nil, "Comma-separated list of assignees")
	createCmd.Flags().StringSliceVarP(&labels, "labels", "l", nil, "Comma-separated list of labels")
	createCmd.Flags().IntVarP(&milestone, "milestone", "m", 0, "Milestone ID")
	createCmd.MarkFlagRequired("title") // Ensure title is required
}
