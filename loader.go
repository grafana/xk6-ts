// Package ts contains xk6-ts extension.
package ts

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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

func isEnabled(filename string) bool {
	if strings.HasSuffix(filename, ".ts") {
		return true
	}

	if !strings.HasSuffix(filename, ".js") {
		return true
	}

	return os.Getenv("XK6_TS") == "always" //nolint:forbidigo
}

//nolint:forbidigo
func redirectStdin() {
	isRun, scriptIndex := isRunCommand(os.Args)
	if !isRun {
		return
	}

	filename := os.Args[scriptIndex]

	if !isEnabled(filename) {
		return
	}

	fmt.Println("Heeee")

	opts := &k6pack.Options{
		Filename:  filename,
		SourceMap: os.Getenv("XK6_TS_SOURCEMAP") == "true",
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
