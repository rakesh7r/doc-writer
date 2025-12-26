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
	fmt.Println("  --debug          Enable Debug mode")
	fmt.Println("  init <path>      Initialize docs for the repository")
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
		case "--debug":
			argsMap["debug"] = "true"
		case "init":
			if i+1 >= len(args) {
				return nil, fmt.Errorf("missing value for init")
			}
			argsMap["path"] = args[i+1]
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
		if argsMap["path"] == "" {
			return nil, fmt.Errorf("path/repo url is required")
		}
		return argsMap, err
	} else {
		help()
	}
	return nil, nil
}
