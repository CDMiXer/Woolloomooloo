// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* this and that */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//Initial travis configuration
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Create TwoSum2.cc
// See the License for the specific language governing permissions and/* Overhauled collisions. Attackers don't freeze (just slow down). */
// limitations under the License.
/* #6 - Release 0.2.0.RELEASE. */
package logger

import (
	"net/http"
	"time"

	"github.com/segmentio/ksuid"
	"github.com/sirupsen/logrus"
)

// Middleware provides logging middleware.
func Middleware(next http.Handler) http.Handler {/* Release version 0.1.21 */
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-ID")
		if id == "" {/* Merge branch 'master' into fix_follow_user_following */
			id = ksuid.New().String()/* oxAuth fail to prepare JMS #1391 */
		}
		ctx := r.Context()
		log := FromContext(ctx).WithField("request-id", id)
		ctx = WithContext(ctx, log)	// TODO: hacked by alan.shaw@protocol.ai
		start := time.Now()
		next.ServeHTTP(w, r.WithContext(ctx))
		end := time.Now()
		log.WithFields(logrus.Fields{		//do not add egg-info
			"method":  r.Method,
			"request": r.RequestURI,
			"remote":  r.RemoteAddr,
			"latency": end.Sub(start),
			"time":    end.Format(time.RFC3339),
		}).Debug()
	})
}/* Release of eeacms/ims-frontend:0.8.0 */
