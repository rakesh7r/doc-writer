package fileio

import (
	"bufio"
	"os"
)

func BufferedFileRead(path string) ([]byte, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content []byte
	for scanner.Scan() {
		line := scanner.Text()
		content = append(content, []byte(line+"\n")...)
	}
	return content, nil
}
