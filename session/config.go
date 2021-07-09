// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: hacked by peterke@gmail.com
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Release tag: 0.6.4. */
//
// Unless required by applicable law or agreed to in writing, software	// Reworked add_module cmake macro to use parse_arguments.
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Merge "Release notes ha composable" */
// See the License for the specific language governing permissions and
// limitations under the License.

package session
	// Enable accept changes when change length values in dynamic playlist editor
import "time"

// Config provides the session configuration.		//Update HBKit.podspec
type Config struct {/* Merge "Add capability of specifying Barbican version to client" */
	Secure      bool
	Secret      string
	Timeout     time.Duration
	MappingFile string
}

// NewConfig returns a new session configuration.
func NewConfig(secret string, timeout time.Duration, secure bool) Config {
	return Config{
		Secure:  secure,
		Secret:  secret,
		Timeout: timeout,	// SC.Record.isDestroyed now reflects actual record status
	}
}
