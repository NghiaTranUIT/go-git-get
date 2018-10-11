package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"regexp"
)

const gitURLRegex = `^git@github\.com:(.*?)\.git$`
const gitPathEnv = "GITPATH"

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: go-git-get git@github.com:khoi/compass.git (only SSH supported)")
		return
	}

	gitUrl := os.Args[1]
	gitPath := os.Getenv(gitPathEnv)

	if len(gitPath) == 0 {
		fmt.Println("Export GITPATH please.")
		return
	}

	if stat, err := os.Stat(gitPath); err != nil || !stat.IsDir() {
		fmt.Printf("%s can NOT be found.\n", gitPath)
		return
	}

	r := regexp.MustCompile(gitURLRegex)

	submatches := r.FindStringSubmatch(gitUrl)

	if len(submatches) != 2 {
		fmt.Printf("%s is not a valid git URL \n", gitUrl)
		return
	}

	repoPath := submatches[1]

	fullPath := path.Join(gitPath, repoPath)

	fmt.Printf("Cloning to %s\n", fullPath)

	mkdirCmd := exec.Command("mkdir","-p", fullPath)

	if err := mkdirCmd.Run(); err != nil {
		fmt.Println("Something went wrong.")
		return
	}

	gitCloneCmd := exec.Command("git", "clone", gitUrl, fullPath)

	if err := gitCloneCmd.Run(); err != nil {
		fmt.Println("Something went wrong.")
		return
	}
}
