# Commit Summary CLI

A command-line tool that generates concise commit messages using OpenAI's GPT model. It can be used standalone or as a git pre-commit hook.

## Installation

```bash
# Install the binary
go install github.com/kodryan/commit-summary@latest

# Or build from source
go build -o summary
mv summary /usr/local/bin/
```

## Configuration

Set your OpenAI API key as an environment variable:
```bash
export OPENAI_API_KEY=your-key-here
```

Or add it to your shell profile (~/.bashrc, ~/.zshrc):
```bash
echo 'export OPENAI_API_KEY=your-key-here' >> ~/.zshrc
```

## Usage

### Command Line
```bash
# Generate commit message from staged changes
summary -diff "$(git diff --cached)"

# Generate message from any diff
summary -diff "your diff content here"
```

### Git Pre-commit Hook

1. Create `.git/hooks/prepare-commit-msg`:
```bash
#!/bin/sh

DIFF_CONTENT=$(git diff --cached)
SUMMARY=$(summary -diff "$DIFF_CONTENT")

COMMIT_MSG_FILE=$1
if [[ -n "$SUMMARY" ]]; then
    echo "$SUMMARY" > "$COMMIT_MSG_FILE"
fi
```

2. Make it executable:
```bash
chmod +x .git/hooks/prepare-commit-msg
```

Now when you run `git commit`, it will automatically generate a commit message based on your staged changes.

## Options

- `-diff string`: The diff content to generate a commit message for (required)

## Error Codes

- Exit 1: Invalid input or API error

