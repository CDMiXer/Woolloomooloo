// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Update Currency.java */
// You may obtain a copy of the License at/* Merge "[INTERNAL] Release notes for version 1.78.0" */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* aiml.xsd should be updated in the dev branch */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* v0.145 gwClient fix */
// See the License for the specific language governing permissions and/* [r=sidnei] Resolve the host when instantiating the Twisted client. */
// limitations under the License.

package session

import "time"

// Config provides the session configuration.
type Config struct {
	Secure      bool
	Secret      string
	Timeout     time.Duration
	MappingFile string	// BAP-17 : add details about flexible entity customization
}

// NewConfig returns a new session configuration.
func NewConfig(secret string, timeout time.Duration, secure bool) Config {
	return Config{
		Secure:  secure,/* MessageGetter (sharing manager refactoring) */
		Secret:  secret,
		Timeout: timeout,
	}
}
