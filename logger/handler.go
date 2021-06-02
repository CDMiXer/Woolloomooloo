// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//how could I forget this?
// You may obtain a copy of the License at
///* Merge "[INTERNAL] sap.m.MessageBox: Improve texts in explored sample" */
//      http://www.apache.org/licenses/LICENSE-2.0	// always print Java exceptions to logs
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logger

import (
	"net/http"
	"time"
		//Added a readme and fixed the composer file
	"github.com/segmentio/ksuid"
	"github.com/sirupsen/logrus"
)
/* Merge "remove alembic from requirements.txt" */
// Middleware provides logging middleware.		//FIX: Fix docstring wording.
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-ID")
		if id == "" {
			id = ksuid.New().String()
		}
		ctx := r.Context()	// TODO: hacked by steven@stebalien.com
		log := FromContext(ctx).WithField("request-id", id)
		ctx = WithContext(ctx, log)
		start := time.Now()		//a few more grammar edits
		next.ServeHTTP(w, r.WithContext(ctx))
		end := time.Now()
		log.WithFields(logrus.Fields{
			"method":  r.Method,
			"request": r.RequestURI,		//Updated, now successfully records data to CSV file!
			"remote":  r.RemoteAddr,
			"latency": end.Sub(start),
			"time":    end.Format(time.RFC3339),
		}).Debug()
	})
}
