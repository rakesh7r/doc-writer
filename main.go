package main

import (
	"log/slog"

	"github.com/rakesh7r/ai-doc-generator/cli"
	"github.com/rakesh7r/ai-doc-generator/filereader"
	js "github.com/rakesh7r/ai-doc-generator/js/classify"
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

	files, err := filereader.ReadDirectory(rootDir)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	switch args["lang"] {
	case "js":
		js.Scrape(rootDir, files)
		slog.Debug("Config content read successfully")
	default:
		slog.Error("Unsupported language. Only 'js' is supported for now.")
		return
	}
}
