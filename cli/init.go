package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func getRootDir(path string) (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	cmd.Dir = path
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func initializeDocs(args map[string]string) error {
	resolved, err := filepath.Abs(args["path"])
	if err != nil {
		panic(err)
	}

	gitDir, err := getRootDir(resolved)
	if err != nil {
		fmt.Printf("%s is not a git repository\n", resolved)
		return err
	}

	gitDir = strings.TrimSpace(gitDir)
	docsDir := filepath.Join(gitDir, ".docs")
	fmt.Printf("Docs initialized at %s\n", docsDir)
	err = os.MkdirAll(docsDir, 0o755)
	if err != nil {
		panic(err)
	}
	return nil
}

func Init(args map[string]string) error {
	return initializeDocs(args)
}
