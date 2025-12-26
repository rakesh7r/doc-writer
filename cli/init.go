package cli

import (
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func getRootDir(path string) (string, error) {
	slog.Debug("Getting root directory", "path", path)
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	cmd.Dir = path
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func isDocsInitialized(docsDir string) (bool, error) {
	_, err := os.Stat(docsDir)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func initializeDocs(args map[string]string) error {
	resolved, err := filepath.Abs(args["path"])
	if err != nil {
		slog.Error("Failed to resolve path", "error", err)
		panic(err)
	}

	gitDir, err := getRootDir(resolved)
	if err != nil {
		slog.Error("Failed to get root directory", "error", err)
		return err
	}

	gitDir = strings.TrimSpace(gitDir)
	docsDir := filepath.Join(gitDir, ".docs")

	isInitialized, err := isDocsInitialized(docsDir)
	if err != nil {
		slog.Error("Failed to check if docs are initialized", "error", err)
		return err
	}
	if isInitialized {
		slog.Info("Docs are already initialized at", "path", docsDir)
		return nil
	}

	err = os.MkdirAll(docsDir, 0o755)
	if err != nil {
		slog.Error("Failed to create docs directory", "error", err)
		panic(err)
	}
	return nil
}

func Init(args map[string]string) error {
	return initializeDocs(args)
}
