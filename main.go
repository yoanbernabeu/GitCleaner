package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// Fun ASCII Art Header
	fmt.Println("\n==============================")
	fmt.Println("Git Cleaner - Simplify your Git history!")
	fmt.Println("Removing files from Git history safely and effectively.")
	fmt.Println("==============================\n")

	// Check if we are in a directory with a git repository
	if !isGitRepository() {
		log.Fatalf("This directory is not a Git repository. Please run the command inside a valid Git repository.")
	}
	// Use the flag package to parse arguments
	filePath := flag.String("file", "", "Path of the file to remove from Git history")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Usage: git-clean --file <file_path>")
		os.Exit(1)
	}

	cleanedPath := filepath.ToSlash(filepath.Clean(*filePath))

	// Check if the file exists
	if _, err := os.Stat(cleanedPath); os.IsNotExist(err) {
		log.Fatalf("The file '%s' does not exist in the current directory.\n", cleanedPath)
	}

	// Add the file to .gitignore if it currently exists
	addFileToGitignore(cleanedPath)

	// Search for commits containing the file
	fmt.Println("Searching for commits containing the file...")
	commits, err := getCommitsWithFile(cleanedPath)
	if err != nil {
		log.Fatalf("Error searching for commits: %v", err)
	}

	if len(commits) == 0 {
		fmt.Printf("The file '%s' was not found in the Git history.\n", cleanedPath)
		os.Exit(0)
	}

	// Display a summary of the commits
	fmt.Printf("The file '%s' is present in %d commit(s):\n", cleanedPath, len(commits))
	for _, commit := range commits {
		fmt.Println(commit)
	}

	// Ask for user confirmation
	if !getUserConfirmation() {
		fmt.Println("Operation canceled by the user.")
		os.Exit(0)
	}

	// Use native git commands to remove the file from history
	removeFileFromHistoryNative(cleanedPath)

	fmt.Println("The file has been removed from the Git history.")
	fmt.Println("Don't forget to force update the remote references with:")
	fmt.Println("git push origin --force --all")
	fmt.Println("git push origin --force --tags")
}

// getCommitsWithFile returns a list of commits where the file is present
func getCommitsWithFile(filePath string) ([]string, error) {
	cmd := exec.Command("git", "log", "--pretty=format:%h %ad | %s%d [%an]", "--date=short", "--", filePath)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	var commits []string
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			commits = append(commits, line)
		}
	}
	return commits, nil
}

// getUserConfirmation asks for user confirmation
func getUserConfirmation() bool {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Do you really want to remove this file from the Git history? (yes/no): ")
		scanner.Scan()
		response := strings.ToLower(strings.TrimSpace(scanner.Text()))
		if response == "yes" || response == "y" {
			return true
		} else if response == "no" || response == "n" {
			return false
		} else {
			fmt.Println("Please answer 'yes' or 'no'.")
		}
	}
}

// removeFileFromHistoryNative uses native git commands to remove the file from history
func removeFileFromHistoryNative(filePath string) {
	// Remove the file from commits
	cmd := exec.Command("git", "filter-branch", "--force", "--index-filter", fmt.Sprintf("git rm --cached --ignore-unmatch '%s'", filePath), "--prune-empty", "--tag-name-filter", "cat", "--", "--all")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Removing the file from the Git history...")

	if err := cmd.Run(); err != nil {
		log.Fatalf("Error removing the file: %v", err)
	}

	// Clean up backup references created by git filter-branch
	if err := os.RemoveAll(".git/refs/original"); err != nil {
		log.Fatalf("Error removing original references: %v", err)
	}
	if err := exec.Command("git", "reflog", "expire", "--expire=now", "--all").Run(); err != nil {
		log.Fatalf("Error expiring reflogs: %v", err)
	}
	if err := exec.Command("git", "gc", "--prune=now", "--aggressive").Run(); err != nil {
		log.Fatalf("Error running garbage collection: %v", err)
	}
}

// isGitRepository checks if the current directory is a Git repository
func isGitRepository() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

// addFileToGitignore adds the file to .gitignore if it is not already there
func addFileToGitignore(filePath string) {
	gitignorePath := ".gitignore"
	file, err := os.OpenFile(gitignorePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("Error opening .gitignore: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == filePath {
			// The file is already in .gitignore
			return
		}
	}

	// Add the file to .gitignore
	if _, err := file.WriteString(filePath + "\n"); err != nil {
		log.Fatalf("Error writing to .gitignore: %v", err)
	}
	fmt.Printf("Added '%s' to .gitignore\n", filePath)
}
