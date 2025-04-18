package internal

import (
	"regexp"
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
