package internal

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func Run(flags Flags) {
	if !checkGitRepository() {
		fmt.Println("This is not a git repository 🙅‍♂️")
		os.Exit(1)
	}

	diff, err := getGitDiff()
	if err != nil {
		fmt.Println("Error getting git diff:", err)
		os.Exit(1)
	}

	originalDiff := diff
	diff = filterLockFiles(diff)

	if strings.TrimSpace(diff) == "" {
		fmt.Println("No changes to commit except lock files 🙅")
		fmt.Println("Maybe you forgot to add files? Try running git add . and then run this script again.")
		os.Exit(1)
	}

	if diff != originalDiff {
		fmt.Println("Changes detected in lock files. These changes will be included in the commit but won't be analyzed for commit message generation.")
	}

	generateCommit(diff, flags)
}

func generateCommit(diff string, flags Flags) {
	prompt := getPromptCommit(diff, flags.CommitType, flags.Language)

	text, err := generateCommitMessage(prompt)
	if err != nil {
		fmt.Println("Error generate commit message:", err)
		os.Exit(1)
	}

	if flags.Emoji {
		text = addGitmojiToCommitMessage(text)
	}

	if flags.Template != "" {
		text = processTemplate(flags.Template, text)
	}

	fmt.Println("Proposed Commit:")
	fmt.Println("------------------------------")
	fmt.Println(text)
	fmt.Println("------------------------------")

	var continueCommit string
	fmt.Print("Do you want to continue? (y/n): ")
	_, err = fmt.Scanln(&continueCommit)
	if err != nil {
		fmt.Println("Commit aborted by user 🙅‍♂️")
		os.Exit(1)
	}

	continueCommit = strings.ToLower(continueCommit)
	if continueCommit != "y" {
		fmt.Println("Commit aborted by user 🙅‍♂️")
		os.Exit(1)
	}

	makeCommit(text)
}

func generateCommitMessage(prompt string) (string, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("GEMINI_API_KEY is missing")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", fmt.Errorf("Error create client Gemini: %v", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.0-flash")

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("Error call Gemini API: %v", err)
	}

	if len(resp.Candidates) > 0 {
		message := resp.Candidates[0].Content.Parts[0].(genai.Text)
		return strings.TrimSpace(string(message)), nil
	}

	return "------------------------------", nil
}
