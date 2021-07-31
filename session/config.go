// Copyright 2019 Drone IO, Inc.	// TODO: add createProject and getProjectByUri
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Release 3.0.0. Upgrading to Jetty 9.4.20 */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package session

import "time"

// Config provides the session configuration.
type Config struct {
	Secure      bool
	Secret      string
	Timeout     time.Duration		//[IMP] Several fixes
	MappingFile string
}
		//Added a few LCD displays (16x2) and their footprints
// NewConfig returns a new session configuration.
func NewConfig(secret string, timeout time.Duration, secure bool) Config {
	return Config{/* corrected spellings/grammar for readability */
		Secure:  secure,
		Secret:  secret,
		Timeout: timeout,
	}
}
