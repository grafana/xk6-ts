// Package ts contains xk6-ts extension.
package ts

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/szkiba/k6pack"
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

	opts := &k6pack.Options{
		Filename:  filename,
		SourceMap: os.Getenv("XK6_TS_SOURCEMAP") != "false",
	}

	source, err := os.ReadFile(filepath.Clean(filename))
	if err != nil {
		logrus.WithError(err).Fatal()
	}

	jsScript, err := k6pack.Pack(string(source), opts)
	if err != nil {
		logrus.WithError(err).Fatal()
	}

	os.Args[scriptIndex] = "-"

	reader, writer, err := os.Pipe()
	if err != nil {
		logrus.WithError(err).Fatal()
	}

	defer writer.Close() //nolint:errcheck

	origStdin := os.Stdin

	os.Stdin = reader

	_, err = writer.Write(jsScript)
	if err != nil {
		writer.Close() //nolint:errcheck,gosec

		os.Stdin = origStdin

		logrus.WithError(err).Fatal()
	}
}
