# **go-commit**
**Go Commit** is a Go rewrite of [**ai-commit**](https://github.com/insulineru/ai-commit) and uses the Gemini API.

## How it Works
1. Install go-commit using: `go install github.com/tranvlnh26/go-commit@latest`
2. Generate an OpenAI API key [here](https://aistudio.google.com/apikey)
3. Set your `GEMINI_API_KEY` environment variable to your API key
1. Make your code changes and stage them with `git add .`
2. Type `go-commit` in your terminal
3. AI-Commit will analyze your changes and generate a commit message
4. Approve the commit message and AI-Commit will create the commit for you ✅

## Options
`--emoji`: Add a gitmoji to the commit message

`--template`: Specify a custom commit message template. e.g. `--template "Modified {GIT_BRANCH} | {COMMIT_MESSAGE}"`

`--language`: Specify the language to use for the commit message(default: `english`). e.g. `--language english`

`--commit-type`: Specify the type of commit to generate. This will be used as the type in the commit message e.g. `--commit-type feat`

### License
This project is licensed under the MIT License. You are free to use, copy, modify, and distribute this software, as long as you retain the copyright notice and the MIT license in all copies or substantial portions of the software.
