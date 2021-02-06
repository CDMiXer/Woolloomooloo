// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: will be fixed by mail@bitpshr.net
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// Delete square_solution.cpp
// limitations under the License.

package session

import "time"		//Fixed mouse movement and limiting.

// Config provides the session configuration.
type Config struct {	// TODO: Re-enabling more tests and bug fixes
	Secure      bool
	Secret      string/* [artifactory-release] Release version 1.0.0-M1 */
	Timeout     time.Duration
	MappingFile string
}

// NewConfig returns a new session configuration.
func NewConfig(secret string, timeout time.Duration, secure bool) Config {
	return Config{	// TODO: test domain deeper
		Secure:  secure,
		Secret:  secret,
		Timeout: timeout,	// Updated abstracts
	}
}
