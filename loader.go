// Package ts contains xk6-ts extension.
package ts

import (
	"os"
	"path/filepath"
	"time"

	"github.com/grafana/k6pack"
	"github.com/sirupsen/logrus"
)

func init() {
	redirectStdin()
}

func isRunCommand(args []string) (bool, int) {
	argn := len(args)

	scriptIndex := argn - 1
	if scriptIndex < 0 {
		return false, scriptIndex
	}

	var runIndex int

	for idx := 0; idx < argn; idx++ {
		arg := args[idx]
		if arg == "run" && runIndex == 0 {
			runIndex = idx

			break
		}
	}

	if runIndex == 0 {
		return false, -1
	}

	return true, scriptIndex
}

//nolint:forbidigo
func redirectStdin() {
	if os.Getenv("XK6_TS") == "false" {
		return
	}

	isRun, scriptIndex := isRunCommand(os.Args)
	if !isRun {
		return
	}

	filename := os.Args[scriptIndex]
	if filename == "-" {
		return
	}

	cwd, _ := os.Getwd()

	opts := &k6pack.Options{
		Filename:   filename,
		SourceMap:  os.Getenv("XK6_TS_SOURCEMAP") != "false",
		SourceRoot: cwd,
	}

	source, err := os.ReadFile(filepath.Clean(filename))
	if err != nil {
		logrus.WithError(err).Fatal()
	}

	packStarted := time.Now()

	jsScript, err := k6pack.Pack(string(source), opts)
	if err != nil {
		logrus.WithError(err).Fatal()
	}

	if os.Getenv("XK6_TS_BENCHMARK") == "true" {
		duration := time.Since(packStarted)
		logrus.WithField("extension", "xk6-ts").WithField("duration", duration).Info("Bundling completed in ", duration)
	}

	os.Args[scriptIndex] = "-"

	reader, writer, err := os.Pipe()
	if err != nil {
		logrus.WithError(err).Fatal()
	}

	os.Stdin = reader

	go func() {
		_, werr := writer.Write(jsScript)
		writer.Close() //nolint:errcheck,gosec

		if werr != nil {
			logrus.WithError(werr).Fatal("stdin redirect failed")
		}
	}()
}
