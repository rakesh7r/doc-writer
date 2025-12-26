package main

import (
	"log/slog"

	"github.com/rakesh7r/ai-doc-generator/cli"
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

	cli.Init(args)

}
