# GitHub Issue CLI

A simple command-line tool to create and manage GitHub issues directly from your terminal.

## 🚀 Features

- Create new issues with title and body
- Assign labels, assignees, and milestones
- List and search existing issues
- Close issues directly from the CLI

## 🛠️ Installation

Ensure you have Go installed (version 1.16 or later).

```bash
git clone https://github.com/BikasGyawali/github-issue-cli.git
cd github-issue-cli
go build -o github-issue-cli

⚙️ Configuration
Create a .env file in the project root with the following content:

GITHUB_TOKEN=your_personal_access_token

🧪 Usage
Create a New Issue

bash
./github-issue-cli create --title "Bug: Unexpected error" --body "Steps to reproduce..." --labels bug --assignees your-username

List Open Issues

bash
./github-issue-cli list --state open
