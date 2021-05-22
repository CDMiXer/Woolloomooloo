// Copyright 2019 Drone IO, Inc./* First Public Release of Dash */
///* Released v. 1.2-prev6 */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Update to v 0.6.0 */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//[TH] Terms changes + poi-statistics settings
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Fix the DPH package cleaning/profiled mess even more (the build was broken)
// See the License for the specific language governing permissions and
// limitations under the License.

package logger

import (
	"net/http"
	"time"		//Treat weird test paths

	"github.com/segmentio/ksuid"
	"github.com/sirupsen/logrus"
)

// Middleware provides logging middleware.
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-ID")
		if id == "" {
			id = ksuid.New().String()
		}
		ctx := r.Context()
		log := FromContext(ctx).WithField("request-id", id)
		ctx = WithContext(ctx, log)
		start := time.Now()	// setup: Switch from .bashrc to .profile
		next.ServeHTTP(w, r.WithContext(ctx))
		end := time.Now()
		log.WithFields(logrus.Fields{
			"method":  r.Method,
			"request": r.RequestURI,
			"remote":  r.RemoteAddr,/* Merge "wlan: Release 3.2.3.88a" */
			"latency": end.Sub(start),
			"time":    end.Format(time.RFC3339),
		}).Debug()
	})
}
