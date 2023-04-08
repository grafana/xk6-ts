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

	"go.k6.io/k6/metrics"
	"go.k6.io/k6/output"
)

type Output struct{}

var _ output.Output = (*Output)(nil)

func New(params output.Params) (output.Output, error) { //nolint:ireturn
	return new(Output), nil
}

func (out *Output) Description() string {
	return "enhanced"
}

func (out *Output) Start() error {
	return ErrUsage
}

func (out *Output) Stop() error {
	return nil
}

func (out *Output) AddMetricSamples(samples []metrics.SampleContainer) {
}

var ErrUsage = errors.New("use --compatibility-mode=enhanced to enable enhanced JavaScript compatibility (TypeScript, etc)")
