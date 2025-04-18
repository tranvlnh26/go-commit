package internal

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func processTemplate(template, commitMessage string) string {
	if !strings.Contains(template, "COMMIT_MESSAGE") {
		fmt.Println("Warning: template doesn't include {COMMIT_MESSAGE}")
		return commitMessage
	}

	finalCommitMessage := strings.ReplaceAll(template, "{COMMIT_MESSAGE}", commitMessage)

	if strings.Contains(finalCommitMessage, "GIT_BRANCH") {
		currentBranchCmd := exec.Command("git", "branch", "--show-current")
		currentBranchBytes, err := currentBranchCmd.Output()
		if err != nil {
			fmt.Println("Error fetching current branch:", err)
			return finalCommitMessage
		}
		currentBranch := strings.TrimSpace(string(currentBranchBytes))

		fmt.Println("Using currentBranch:", currentBranch)

		finalCommitMessage = strings.ReplaceAll(finalCommitMessage, "{GIT_BRANCH}", currentBranch)
	}

	return strings.TrimSpace(finalCommitMessage)
}

func getPromptCommit(diff string, commitType, language string) string {
	prompt := fmt.Sprintf(
		"Write a professional git commit message based on the diff below in %s language", language,
	)

	if commitType != "" {
		prompt += fmt.Sprintf(" with commit type '%s' (lowercase). ", commitType)
	} else {
		prompt += ". "
	}

	prompt += " Do not preface the commit with anything, use the present tense, return the full sentence and also commit type (lowercase):"

	prompt += "\n\n" + diff

	return prompt
}

func getGitDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--cached")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func checkGitRepository() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	output, err := cmd.Output()
	if err != nil {
		return false
	}

	return bytes.TrimSpace(output) != nil && string(bytes.TrimSpace(output)) == "true"
}

func filterLockFiles(diff string) string {
	lines := strings.Split(diff, "\n")
	isLockFile := false
	var filteredLines []string

	for _, line := range lines {
		if strings.HasPrefix(line, "diff --git a/") && (strings.Contains(line, "yarn.lock") || strings.Contains(line, "pnpm-lock.yaml") || strings.Contains(line, "package-lock.json")) {
			isLockFile = true
			continue
		}
		if isLockFile && strings.HasPrefix(line, "diff --git") {
			isLockFile = false
		}
		if !isLockFile {
			filteredLines = append(filteredLines, line)
		}
	}

	return strings.Join(filteredLines, "\n")
}

func makeCommit(input string) {
	trimmedInput := strings.TrimSpace(input)

	fmt.Println("Committing Message... 🚀 ")

	cmd := exec.Command("git", "commit", "-F", "-")
	cmd.Stdin = strings.NewReader(trimmedInput) // Pass the commit message as stdin

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error committing message: %v\n", err)
		return
	}

	fmt.Println("Commit Successful! 🎉")
}
