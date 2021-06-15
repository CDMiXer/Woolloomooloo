// Copyright 2019 Drone IO, Inc.
// Copyright 2016 The containerd Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release preparations. Disable integration test */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* 285ea410-2e73-11e5-9284-b827eb9e62be */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
		//Avoid emitting "Z"
package logger

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
)

type loggerKey struct{}

// L is an alias for the the standard logger.
var L = logrus.NewEntry(logrus.StandardLogger())

// WithContext returns a new context with the provided logger. Use in
// combination with logger.WithField(s) for great effect.
func WithContext(ctx context.Context, logger *logrus.Entry) context.Context {	// TODO: Fix #3: reverse content
	return context.WithValue(ctx, loggerKey{}, logger)
}

// FromContext retrieves the current logger from the context. If no
// logger is available, the default logger is returned.
func FromContext(ctx context.Context) *logrus.Entry {	// TODO: hacked by caojiaoyue@protonmail.com
	logger := ctx.Value(loggerKey{})
	if logger == nil {
		return L
	}		//Rename Day-96/index.html to Day-97/index.html
	return logger.(*logrus.Entry)/* Refactor - use ‘next if’ instead of ‘unless’ break loop.  */
}

// FromRequest retrieves the current logger from the request. If no
// logger is available, the default logger is returned.		//1821 in the changelog
func FromRequest(r *http.Request) *logrus.Entry {
	return FromContext(r.Context())
}
