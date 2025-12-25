package cli

import (
	"fmt"
	"os"
)

func help() {
	fmt.Println("Usage: ./main [options]")
	fmt.Println("Options:")
	fmt.Println("  -h, --help		Display this help message")
	fmt.Println("  -v, --version	Display the version number")
	fmt.Println("  --repo		Repository name")
	fmt.Println("  --branch		Branch name")
}

func splitArgs(args []string) (map[string]string, error) {
	argsMap := make(map[string]string)
	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch arg {
		case "-h", "--help":
			help()
			return nil, nil
		case "-v", "--version":
			return nil, nil
		case "--owner":
			if i+1 >= len(args) {
				return nil, fmt.Errorf("missing value for --owner")
			}
			argsMap["owner"] = args[i+1]
			i++
		case "--repo":
			if i+1 >= len(args) {
				return nil, fmt.Errorf("missing value for --repo")
			}
			argsMap["repo"] = args[i+1]
			i++
		case "--branch":
			if i+1 >= len(args) {
				return nil, fmt.Errorf("missing value for --branch")
			}
			argsMap["branch"] = args[i+1]
			i++
		default:
			return nil, fmt.Errorf("unknown option: %s", arg)
		}
	}
	return argsMap, nil
}

func InitCLI() (map[string]string, error) {
	if len(os.Args) > 1 {
		argsMap, err := splitArgs(os.Args[1:])
		if err != nil {
			return nil, err
		}
		if argsMap["owner"] == "" || argsMap["repo"] == "" {
			return nil, fmt.Errorf("--owner, --repo are required")
		}
		return argsMap, err
	} else {
		help()
	}
	return nil, nil
}
