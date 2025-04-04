package internal

import (
	"regexp"
	"strings"
)

func addGitmojiToCommitMessage(commitMessage string) string {
	typeToGitmoji := map[string]string{
		"feat":     "✨",
		"fix":      "🚑",
		"docs":     "📝",
		"style":    "💄",
		"refactor": "♻️",
		"test":     "✅",
		"chore":    "🔧",
	}
	commitMessage = strings.ToLower(commitMessage)

	re := regexp.MustCompile(`^[a-zA-Z]+`)
	match := re.FindString(commitMessage)

	if match == "" {
		return commitMessage
	}

	if emoji, ok := typeToGitmoji[match]; ok {
		return emoji + " " + commitMessage
	}

	return commitMessage
}
