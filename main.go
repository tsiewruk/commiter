package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	maxLength   = 72
	totalLength = 50 // Szerokość tabeli dla wyśrodkowania nagłówka
	header      = "Commit Message Validation:"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: commit-msg <commit-message-file>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	commitMessage, err := readCommitMessage(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(1)
	}

	status := validateCommitMessage(commitMessage)

	printResult(commitMessage, status)

	if strings.Contains(status, "❌") {
		os.Exit(1)
	}

	os.Exit(0)
}

func readCommitMessage(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var commitMessage strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		commitMessage.WriteString(scanner.Text() + "\n")
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return strings.TrimSpace(commitMessage.String()), nil
}

func validateCommitMessage(commitMessage string) string {
	commitRegex := `^(feat|fix|docs|style|refactor|perf|test|build|ci|chore|revert)(\([a-z ]+\))?: .{1,100}$`
	re := regexp.MustCompile(commitRegex)

	var status strings.Builder
	if re.MatchString(commitMessage) {
		status.WriteString("✅ Message format is valid.\n")
	} else {
		status.WriteString("❌ Invalid commit message format.\n")
	}

	lineTooLong := false
	lines := strings.Split(commitMessage, "\n")
	for _, line := range lines {
		lineLength := len(line)
		if lineLength > maxLength {
			lineTooLong = true
			status.WriteString(fmt.Sprintf("❌ Line exceeds %d characters (Actual length: %d).\n", maxLength, lineLength))
		}
	}

	if !lineTooLong {
		status.WriteString("✅ All lines are within the allowed length.\n")
	}

	return status.String()
}

func printResult(commitMessage, status string) {
	padding := (totalLength - len(header)) / 2
	headerLine := fmt.Sprintf("%s%s%s", strings.Repeat(" ", padding), header, strings.Repeat(" ", padding))

	fmt.Println(strings.Repeat(" ", padding) + "Commit Message:")
	fmt.Println(strings.Repeat("-", totalLength))
	fmt.Println(commitMessage + "\n")
	fmt.Println(strings.Repeat("-", totalLength))
	fmt.Println(headerLine)
	fmt.Println(strings.Repeat("-", totalLength))
	fmt.Print(status)
	fmt.Println(strings.Repeat("-", totalLength))
}
