package main

import (
	"fmt"

	"github.com/rakesh7r/ai-doc-generator/cli"
)

func main() {
	args, err := cli.InitCLI()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(args)
}
