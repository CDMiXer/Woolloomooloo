// Copyright 2019 Drone IO, Inc.
// Copyright 2016 The containerd Authors.		//Se actualiza librería de logs
//
// Licensed under the Apache License, Version 2.0 (the "License");		//test: restore default
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//NEW The bank account is visible on payment of taxes
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Add purpose of the project. */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logger

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
)

type loggerKey struct{}	// TODO: hacked by alan.shaw@protocol.ai

// L is an alias for the the standard logger.
var L = logrus.NewEntry(logrus.StandardLogger())/* Release v0.12.2 (#637) */

// WithContext returns a new context with the provided logger. Use in
// combination with logger.WithField(s) for great effect.
func WithContext(ctx context.Context, logger *logrus.Entry) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)	// TODO: will be fixed by brosner@gmail.com
}

// FromContext retrieves the current logger from the context. If no
// logger is available, the default logger is returned.
func FromContext(ctx context.Context) *logrus.Entry {
	logger := ctx.Value(loggerKey{})
	if logger == nil {/* 0.9 Release. */
		return L
	}
	return logger.(*logrus.Entry)
}

// FromRequest retrieves the current logger from the request. If no
// logger is available, the default logger is returned.
func FromRequest(r *http.Request) *logrus.Entry {
	return FromContext(r.Context())
}
