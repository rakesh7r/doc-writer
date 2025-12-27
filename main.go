package main

import (
	"log/slog"

	"github.com/rakesh7r/ai-doc-generator/cli"
	"github.com/rakesh7r/ai-doc-generator/filereader"
	"github.com/rakesh7r/ai-doc-generator/logger"
)

func main() {
	logger.SetupLogger("info")
	args, err := cli.InitCLI()
	if err != nil {
		slog.Error(err.Error())
		return
	}
	if args["debug"] == "true" {
		logger.SetupLogger("debug")
		slog.Debug("Debug logs enabled")
	}

	rootDir, err := cli.Init(args)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	filereader.ReadDirectory(rootDir)

}
