package main

import (
	"context"
	"fmt"

	"github.com/rakesh7r/ai-doc-generator/cli"
	"github.com/rakesh7r/ai-doc-generator/config"
	ghclient "github.com/rakesh7r/ai-doc-generator/ghclient"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	args, err := cli.InitCLI()
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx := ghclient.WithGitHubToken(context.Background(), config.GitHubToken)

	metadata, err := ghclient.NewGitHubClient(ctx).GetRepo(args["owner"], args["repo"])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("metadata", metadata)
}
