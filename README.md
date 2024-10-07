# Git Cleaner - Simplify Your Git History!

Git Cleaner is a command-line tool designed to help you easily remove files from your Git history. It allows you to delete files that should no longer be tracked, even if they are present in previous commits. The tool provides an interactive experience to ensure safe removal, making your repository cleaner and smaller.

## Features

- File History Removal: Search for and remove a specific file from the entire Git history.

- Automatic .gitignore Update: If the specified file exists in the current directory, it is automatically added to the .gitignore file to prevent future tracking.

- Interactive Confirmation: Git Cleaner will prompt for user confirmation before making any destructive changes.

- Git Filter-Repo vs Native Git: Git Cleaner supports by default git-filter-repo, but if it is not available, it will use native Git commands to remove the file from the history.

## Installation

>  **Recommended**: Install Git Filter-Repo to ensure the best performance: [Git Filter-Repo Installation Guide](https://github.com/newren/git-filter-repo/blob/main/INSTALL.md)

To install Git Cleaner, run the following command:

```bash
curl -sSL https://raw.githubusercontent.com/yoanbernabeu/GitCleaner/main/install.sh | bash
```

## Usage

To use Git Cleaner, run the following command (replace `<file_path>` with the path of the file you want to remove from the Git history):

```bash
git-cleaner --file <file_path>
```

Replace `<file_path>` with the path of the file you want to remove from the Git history.

### Example

```bash
git-cleaner --file secrets.txt
```

This command will search for all the commits containing `secrets.txt` and then prompt you to confirm its removal from the Git history.

## How It Works

1. Search for Commits: It searches for all the commits containing the file.

2. User Confirmation: It provides a list of commits and asks for user confirmation to proceed with removing the file from the history.

3. Remove File from History: If confirmed, it removes the file from the Git history using either git-filter-repo or native Git commands.

4. Add to .gitignore: If the file is present, it will be added to .gitignore to ensure it won't be tracked in future commits.

## Important Notes

- **Force Push Required**: After running Git Cleaner, you need to force push to update the remote repository:
  
    ```bash
    git push origin --force --all
    git push origin --force --tags
    ```

- **Destructive Operation**: Removing files from Git history is a destructive operation. It rewrites the commit history, so be sure all collaborators are aware and understand the implications.

## Requirements

- **Git**: Git must be installed and accessible from the command line.

- **git-filter-repo** (Recommended): If git-filter-repo is not available, Git Cleaner will use native Git commands to remove the file from the history.

- **Go**: The Go runtime is required if you want to build and run the tool from source.

## Disclaimer

Git Cleaner rewrites the Git commit history, which can be risky if not done properly. It is recommended to make backups before running this tool and to coordinate with your team if you are working in a shared repository.

## License

This project is open-source and available under the MIT License.