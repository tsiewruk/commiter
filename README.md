# Commiter
A script that allows you to check the length of a commit and its validation against conventional commit.

## About the script
- The commit length limit is checked; if it is greater than 72 characters, an error is displayed.
- Compliance with the conventional commit is checked; if this format is not followed, an error is returned.

## Installation
```
curl https://github.com/tsiewruk/commiter/blob/main/commit-msg -o </path_to_your_git_project/.git/hooks/commit-msg>
```