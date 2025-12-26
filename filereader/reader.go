package filereader

import (
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

func findGitIgnoreFiles(path string) ([]string, error) {
	// There can be multiple .gitignore files in a repo (e.g., in subdirectories)
	var gitIgnoreFiles []string
	err := filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && d.Name() == ".gitignore" {
			gitIgnoreFiles = append(gitIgnoreFiles, p)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return gitIgnoreFiles, nil
}

func readGitIgnoreFiles(path string) []string {
	ignore := []string{}
	file, err := filepath.Abs(filepath.Join(path, ".gitignore"))
	if err != nil {
		log.Fatal(err)
	}
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	lines := string(data)
	for _, line := range filepath.SplitList(lines) {
		patterns := strings.Split(line, "\n")
		ignore = append(ignore, patterns...)
	}
	ignore = append(ignore, ".git")
	return ignore
}

func ReadDirectory(basePath string) ([]string, error) {
	gitIgnorePaths, _ := findGitIgnoreFiles(basePath)
	ignorePatterns := []string{}

	for _, gitIgnorePath := range gitIgnorePaths {
		patterns := readGitIgnoreFiles(filepath.Dir(gitIgnorePath))
		ignorePatterns = append(ignorePatterns, patterns...)
	}

	files := []string{}

	err := filepath.WalkDir(basePath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err // Propagate errors
		}
		if d.IsDir() {
			for _, pattern := range ignorePatterns {
				matched, err := filepath.Match(pattern, d.Name())
				if err != nil {
					return err
				}
				if matched {
					return filepath.SkipDir // Skip this directory
				}
			}
		} else {
			// It's a file
			for _, pattern := range ignorePatterns {
				matched, err := filepath.Match(pattern, d.Name())
				if err != nil {
					return err
				}
				if matched {
					return nil // Skip this file
				}
			}
			relativePath, err := filepath.Rel(basePath, path)
			if err != nil {
				return err
			}
			files = append(files, relativePath)
		}

		return nil
	})

	slog.Info("no. of files found", "count", len(files))
	if err != nil {
		log.Fatal(err)
	}
	return files, nil
}
