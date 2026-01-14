package js

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/rakesh7r/ai-doc-generator/fileio"
)

type PackageFileName string

var PackageFile PackageFileName = "package.json"

type PackageJson struct {
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	Version      string            `json:"version"`
	Author       string            `json:"author"`
	License      string            `json:"license"`
	Repository   string            `json:"repository"`
	Keywords     []string          `json:"keywords"`
	Dependencies map[string]string `json:"dependencies"`
	DevDeps      map[string]string `json:"dev_dependencies"`
	Scripts      map[string]string `json:"scripts"`

	Timeout int `json:"timeout_seconds"`
}

func ConfigContent(baseUrl string, files map[string][]string) ([]PackageJson, error) {

	var content []PackageJson
	paths, exists := files[string(PackageFile)]
	if exists {
		for _, path := range paths {
			file, err := os.Open(filepath.Join(baseUrl, path))
			if err != nil {
				return nil, err
			}
			var pkgJson PackageJson
			decoder := json.NewDecoder(file)
			if err := decoder.Decode(&pkgJson); err != nil {
				return content, err
			}

			content = append(content, pkgJson)
		}
	} else {
		return nil, os.ErrNotExist
	}

	return content, nil
}

func MarkdownContentIfExists(baseUrl string, files map[string][]string) []string {
	var content []byte
	var readmes []string

	markdownFilePattern := "*.md"
	var markdownPaths []string
	for key := range files {
		matched, err := filepath.Match(markdownFilePattern, filepath.Base(key))
		if matched {
			markdownPaths = append(markdownPaths, files[key]...)
		}
		if err != nil {
			return nil
		}
	}

	for _, path := range markdownPaths {
		file, err := fileio.BufferedFileRead(filepath.Join(baseUrl, path))
		if err != nil {
			return nil
		}
		content = append(content, file...)
		readmes = append(readmes, string(content))
	}
	return readmes
}

func GetDockerContentsIfExists(baseUrl string, files map[string][]string) []string {
	possibleFiles := []string{"Dockerfile", "docker-compose.yml", "docker-compose.yaml"}

	var dockerContents []string
	for _, fileName := range possibleFiles {
		paths, exists := files[fileName]
		if exists {
			for _, path := range paths {
				fileContent, err := fileio.BufferedFileRead(filepath.Join(baseUrl, path))
				if err != nil {
					continue
				}
				dockerContents = append(dockerContents, string(fileContent))
			}
		}
	}
	return dockerContents
}

func Scrape(baseUrl string, files map[string][]string) {
	for _, paths := range files {
		fmt.Println(filepath.Join(baseUrl, paths[0]))
	}
	_, err := ConfigContent(baseUrl, files)
	if err != nil {
		slog.Error("Error reading Package.json content", "error", err)
		return
	}
	GetDockerContentsIfExists(baseUrl, files)
	MarkdownContentIfExists(baseUrl, files)

}
