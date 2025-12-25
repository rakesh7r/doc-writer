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
		if argsMap["repo"] == "" || argsMap["branch"] == "" {
			return nil, fmt.Errorf("--repo and --branch are required")
		}
		return argsMap, err
	} else {
		help()
	}

	// // --- Method 2: Using the 'flag' package (Parsed Flags) ---
	// fmt.Println("\n--- flag package ---")

	// // Define flags (returns pointers)
	// name := flag.String("name", "Guest", "Name of the user")
	// age := flag.Int("age", 0, "Age of the user")
	// isVerbose := flag.Bool("v", false, "Enable verbose logging")

	// // IMPORTANT: Must call flag.Parse() to populate the variables
	// flag.Parse()

	// // Access flags using point dereferencing
	// fmt.Printf("Name: %s\n", *name)
	// fmt.Printf("Age: %d\n", *age)
	// fmt.Printf("Verbose: %t\n", *isVerbose)

	// // flag.Args() returns any non-flag arguments left over after parsing
	// fmt.Printf("Trailing arguments: %v\n", flag.Args())
	return nil, nil
}
