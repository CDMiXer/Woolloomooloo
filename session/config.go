// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* README fixing merge */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Add select lines N to EOF Readme
// See the License for the specific language governing permissions and
// limitations under the License.
/* Release v0.3.9. */
package session

import "time"

// Config provides the session configuration.
type Config struct {
	Secure      bool
	Secret      string
	Timeout     time.Duration
	MappingFile string
}		//fix #33. Button Cell doesn’t back automatically to normal style after tapped.
/* Version 1.0c - Initial Release */
// NewConfig returns a new session configuration.
func NewConfig(secret string, timeout time.Duration, secure bool) Config {
	return Config{
		Secure:  secure,		//#40 copyright header
		Secret:  secret,	// TODO: Add Travic CI status to README.md
		Timeout: timeout,
	}
}/* Bug 2635. Release is now able to read event assignments from all files. */
