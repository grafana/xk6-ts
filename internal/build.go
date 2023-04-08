// MIT License
//
// Copyright (c) 2023 Iván Szkiba
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package internal

import (
	"errors"
	"os"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/sirupsen/logrus"
)

var (
	ErrBuild = errors.New("build error")
	ErrStdin = errors.New("stdin not supported in enhanced mode")
)

func checkArgs(args []string) bool {
	argn := len(args)

	var runIndex, modeIndex int

	for idx := 0; idx < argn; idx++ {
		arg := args[idx]
		if arg == "run" && runIndex == 0 {
			runIndex = idx

			continue
		}

		if runIndex != 0 && arg == "--compatibility-mode=enhanced" {
			modeIndex = idx
			args[idx] = "--compatibility-mode=extended"

			continue
		}
	}

	if os.Getenv("K6_COMPATIBILITY_MODE") == "enhanced" {
		os.Setenv("K6_COMPATIBILITY_MODE", "extended")

		modeIndex++
	}

	return modeIndex != 0
}

func Build() {
	if !checkArgs(os.Args) {
		return
	}

	scriptIndex := len(os.Args) - 1

	script := os.Args[scriptIndex]
	if script == "-" {
		logrus.WithError(ErrStdin).Fatal()
	}

	os.Args[scriptIndex] = "-"

	content, err := load(script)
	if err != nil {
		logrus.WithError(err).Fatal()
	}

	reader, writer, err := os.Pipe()
	if err != nil {
		logrus.WithError(err).Fatal()
	}

	origStdin := os.Stdin

	os.Stdin = reader

	_, err = writer.Write(content)
	if err != nil {
		writer.Close()

		os.Stdin = origStdin

		logrus.WithError(err).Fatal()
	}
}

func load(filename string) ([]byte, error) {
	result := api.Build(api.BuildOptions{ //nolint:exhaustruct
		EntryPoints: []string{filename},
		Bundle:      true,
		Write:       false,
		LogLevel:    api.LogLevelWarning,
		Target:      api.ES5,
		Platform:    api.PlatformBrowser,
		Format:      api.FormatESModule,
		External:    []string{"k6"},
	})

	if len(result.Errors) > 0 {
		return nil, ErrBuild
	}

	return result.OutputFiles[0].Contents, nil
}
